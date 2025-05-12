package models

import "time"

type JobHealth struct {
	JobId         JobId
	Healthy       bool
	LastCheckedAt time.Time
	Message       string
}

type JobsHealthResult struct {
	Healthy   []*Job
	Unhealthy []*Job
	Statuses  map[JobId]JobHealth
}

type ProcessEvent struct {
	JobId     JobId
	EventType ProcessEventType
	Timestamp time.Time
	ExitCode  int
	Error     error
}

type ProcessEventType string

const (
	ProcessStarted ProcessEventType = "started"
	ProcessExited  ProcessEventType = "exited"
	ProcessFailed  ProcessEventType = "failed"
)
