package models

import (
	"fmt"
	"time"
)

type ProcessEventType string

const (
	ProcessStarted ProcessEventType = "STARTED"
	ProcessExited  ProcessEventType = "EXITED"
	ProcessFailed  ProcessEventType = "FAILED"
)

var ProcessEventStateMapping = map[ProcessEventType]JobStatus{
	ProcessStarted: StatusRunning,
	ProcessExited:  StatusExited,
	ProcessFailed:  StatusFatal,
}

type ProcessEvent struct {
	JobId     JobId
	EventType ProcessEventType
	Timestamp time.Time
	ExitCode  int
	Pid       int
	Error     error
}

func NewProcessStartedEvent(jobId JobId, pid int) ProcessEvent {
	return ProcessEvent{
		JobId:     jobId,
		EventType: ProcessStarted,
		Timestamp: time.Now(),
		Pid:       pid,
	}
}

func NewProcessExitedEvent(jobId JobId, pid int, exitCode int) ProcessEvent {
	return ProcessEvent{
		JobId:     jobId,
		EventType: ProcessExited,
		Timestamp: time.Now(),
		ExitCode:  exitCode,
		Pid:       pid,
	}
}

func NewProcessFailedEvent(jobId JobId, pid int, err error) ProcessEvent {
	return ProcessEvent{
		JobId:     jobId,
		EventType: ProcessFailed,
		Timestamp: time.Now(),
		Error:     err,
		Pid:       pid,
	}
}

func (e ProcessEvent) Apply(job *Job) error {
	return job.HandleProcessEvent(e)
}

func (e ProcessEvent) GetTargetState() JobStatus {
	return ProcessEventStateMapping[e.EventType]
}

func (e ProcessEvent) GetStateDescription() string {
	switch e.EventType {
	case ProcessStarted:
		return fmt.Sprintf("Process started with PID %d", e.Pid)
	case ProcessExited:
		return fmt.Sprintf("Process exited with code %d", e.ExitCode)
	case ProcessFailed:
		if e.Error != nil {
			return fmt.Sprintf("Process failed: %v", e.Error)
		}
		return "Process failed with unknown error"
	default:
		return fmt.Sprintf("Unknown process event: %s", e.EventType)
	}
}

func (e ProcessEvent) ToJobState() *JobState {
	return &JobState{
		Status:      e.GetTargetState(),
		ExitCode:    e.ExitCode,
		Description: e.GetStateDescription(),
	}
}

func (e ProcessEvent) CanApplyToStateMachine(sm *StateMachine) bool {
	return sm.CanTransition(e.GetTargetState())
}

func (e ProcessEvent) ApplyToStateMachine(sm *StateMachine) error {
	return sm.Transition(
		e.GetTargetState(),
		e.ExitCode,
		e.GetStateDescription(),
	)
}
