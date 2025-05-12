package interfaces

import (
	models "github.com/touline-p/task-master/supervisor/domain/models"
)

// Monitors the health of running processes
type IWatcherService interface {
	Start() error
	Stop() error

	CheckHealth() *models.JobsHealthResult
	RegisterJob(job *models.Job) error
	UnregisterJob(jobId models.JobId) error
}
