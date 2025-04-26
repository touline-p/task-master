package models

type JobsHealthResult interface {
	HealthyJobs() []Job
	UnHealthyJobs() []Job
}
