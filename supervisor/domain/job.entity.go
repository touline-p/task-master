package supervisor

import (
	"time"
)

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

type Job interface {
	Id() string
	Status() JobStatus
	ConfigValues() []ConfigValue
	CreatedAt() Time
}

func (j *Job) markStarted() {
	j.Status = JobStatus.STARTED
}

func (j *Job) getPossibleTransition() {

}
