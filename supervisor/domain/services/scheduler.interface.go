package services

import (
	"github.com/touline-p/task-master/supervisor/domain/models"
)

// Internal job scheduling and lifecycle
type ISchedulerService interface {
	// Shutdown() error

	RegisterJobs(jobs []models.Job) error
	// UnregisterJobs(jobIds []models.JobId) error
	// UpdateJobs(jobs []models.Job) error

	// TerminateJob(jobId models.JobId) error

	// HandleProcessEvent(event models.ProcessEvent) error
}
