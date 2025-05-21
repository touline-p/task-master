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
