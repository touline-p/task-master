package models

import "slices"

type JobStatus string

const (
	StatusInitial  JobStatus = "initial"
	StatusStarting JobStatus = "starting"
	StatusRunning  JobStatus = "running"
	StatusStopping JobStatus = "stopping"
	StatusStopped  JobStatus = "stopped"
	StatusExited   JobStatus = "exited"
	StatusBackoff  JobStatus = "backoff"
	StatusFatal    JobStatus = "fatal"
)

type JobState struct {
	Status      JobStatus
	ExitCode    int
	Description string
}

type StateMachine struct {
	Current *JobState
}

type StateTransitionError struct {
	Message    string
	FromStatus JobStatus
	ToStatus   JobStatus
}

func (e StateTransitionError) Error() string {
	return e.Message
}

var allowedTransitions = map[JobStatus][]JobStatus{
	StatusInitial:  {StatusStarting},
	StatusStarting: {StatusRunning, StatusBackoff, StatusFatal},
	StatusRunning:  {StatusStopping, StatusExited},
	StatusStopping: {StatusStopped, StatusFatal},
	StatusStopped:  {StatusStarting},
	StatusExited:   {StatusStarting, StatusBackoff, StatusFatal},
	StatusBackoff:  {StatusStarting, StatusFatal},
	StatusFatal:    {StatusInitial},
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		Current: &JobState{
			Status:      StatusInitial,
			Description: "Job initialized",
		},
	}
}

func (sm *StateMachine) CanTransition(to JobStatus) bool {
	if sm.Current == nil {
		return false
	}

	return slices.Contains(allowedTransitions[sm.Current.Status], to)
}

func (sm *StateMachine) Transition(to JobStatus, exitCode int, description string) error {
	if !sm.CanTransition(to) {
		return StateTransitionError{
			Message:    "Invalid state transition",
			FromStatus: sm.Current.Status,
			ToStatus:   to,
		}
	}

	sm.Current = &JobState{
		Status:      to,
		ExitCode:    exitCode,
		Description: description,
	}

	return nil
}
