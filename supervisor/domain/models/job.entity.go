package models

import (
	"fmt"
	"os"
	"slices"
	"time"
)

type JobId string

type RestartPolicy string

const (
	RestartAlways     RestartPolicy = "always"
	RestartNever      RestartPolicy = "never"
	RestartUnexpected RestartPolicy = "unexpected"
)

type JobConfig struct {
	Name          string
	Command       string
	NumProcs      int
	Umask         os.FileMode
	WorkingDir    string
	AutoStart     bool
	AutoRestart   RestartPolicy
	ExitCodes     []int
	StartRetries  int
	StartDuration time.Duration
	StopSignal    string
	StopDuration  time.Duration
	Stdout        string
	Stderr        string
	Environment   map[string]string
	ConfigValues  []JobConfigValue
}

type Job struct {
	Id           JobId
	Config       JobConfig
	StateMachine *StateMachine
	Pid          int
	StartedAt    time.Time
	StopAttempts int
	RetryCount   int
}

func NewJob(id JobId, config JobConfig) *Job {
	fmt.Printf("NewJob %s, config : %+v\n", id, config)

	return &Job{
		Id:           id,
		Config:       config,
		StateMachine: NewStateMachine(id),
		RetryCount:   0,
		StopAttempts: 0,
	}
}

func (j *Job) GetConfigValue(key string) (JobConfigValue, bool) {
	for _, cv := range j.Config.ConfigValues {
		if cv.Key() == key {
			return cv, true
		}
	}
	return JobConfigValue{}, false
}

func (j *Job) AddConfigValue(cv JobConfigValue) {
	j.Config.ConfigValues = append(j.Config.ConfigValues, cv)
}

func (j *Job) GetState() *JobState {
	return j.StateMachine.Current
}

func (j *Job) IsExpectedExit(exitCode int) bool {
	return slices.Contains(j.Config.ExitCodes, exitCode)
}

func (j *Job) HasExceededRetries() bool {
	return j.RetryCount >= j.Config.StartRetries
}

func (j *Job) IsRunningLongEnough() bool {
	if j.StartedAt.IsZero() {
		return false
	}
	return time.Since(j.StartedAt) >= j.Config.StartDuration
}

func (j *Job) ShouldRestart(exitCode int) bool {
	switch j.Config.AutoRestart {
	case RestartAlways:
		return true
	case RestartNever:
		return false
	case RestartUnexpected:
		return !j.IsExpectedExit(exitCode)
	default:
		return false
	}
}

func (j *Job) Start() error {
	err := j.StateMachine.Transition(StatusStarting, 0, "Job is starting")
	j.RetryCount = 0
	return err
}

func (j *Job) MarkAsRunning(pid int) error {
	err := j.StateMachine.Transition(StatusRunning, 0, "Job is running")
	if err == nil {
		j.Pid = pid
		j.StartedAt = time.Now()
	}
	return err
}

func (j *Job) Stop() error {
	return j.StateMachine.Transition(StatusStopping, 0, "Job is stopping")
}

func (j *Job) MarkAsStopped() error {
	j.Pid = 0
	j.StopAttempts = 0
	return j.StateMachine.Transition(StatusStopped, 0, "Job stopped successfully")
}

func (j *Job) MarkAsExited(exitCode int) error {
	j.Pid = 0
	return j.StateMachine.Transition(StatusExited, exitCode, "Job exited")
}

func (j *Job) BackOff() error {
	j.RetryCount++
	j.Pid = 0
	return j.StateMachine.Transition(StatusBackoff, 0, "Job backed off for retry")
}

func (j *Job) MarkAsFatal(reason string) error {
	j.Pid = 0
	return j.StateMachine.Transition(StatusFatal, 0, reason)
}

func (j *Job) IsAlive() bool {
	status := j.StateMachine.Current.Status
	return status == StatusStarting || status == StatusRunning
}

func (j *Job) IsStartable() bool {
	return j.StateMachine.CanTransition(StatusStarting)
}

func (j *Job) IsStoppable() bool {
	return j.StateMachine.CanTransition(StatusStopping)
}

func (j *Job) HandleProcessEvent(event ProcessEvent) error {
	if !event.CanApplyToStateMachine(j.StateMachine) {
		return fmt.Errorf("job %s cannot handle event %s in current state %s",
			j.Id, event.EventType, j.StateMachine.Current.Status)
	}

	switch event.EventType {
	case ProcessStarted:
		j.Pid = event.Pid
		j.StartedAt = event.Timestamp
		return event.ApplyToStateMachine(j.StateMachine)

	case ProcessExited:
		if err := event.ApplyToStateMachine(j.StateMachine); err != nil {
			return err
		}
		j.Pid = 0
		if j.ShouldRestart(event.ExitCode) {
			if j.HasExceededRetries() {
				return j.MarkAsFatal("Exceeded maximum number of retries")
			}
			return j.BackOff()
		}
		return nil

	case ProcessFailed:
		j.Pid = 0
		return event.ApplyToStateMachine(j.StateMachine)

	default:
		return fmt.Errorf("unknown process event type: %s", event.EventType)
	}
}
