package models

type JobStatus int

const (
	STARTING JobStatus = iota
	STARTED
	RUNNING
	STOPPING
	STOPPED
	EXITED
	BACKOFF
	FATAL
)

type JobState interface {
}
