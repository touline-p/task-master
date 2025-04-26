package models

import "time"

type JobId string

type Job interface {
	Id() JobId
	CreatedAt() time.Time
	ConfigValues() []ConfigValue
	State() JobState

	transitionTo(JobState)

	markStarting()
	markStarted()
	markRunning()
	markStopping()
	markStopped()
	markExited()
	markBackoff()
	markFatal()
}
