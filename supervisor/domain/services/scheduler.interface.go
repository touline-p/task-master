package interfaces

import (
	"github.com/touline-p/task-master/supervisor/domain/models"
)

// Internal job scheduling and lifecycle
type ISchedulerService interface {
	Initialize() error
	Shutdown() error

	RegisterJobs(jobs []models.Job) error
	UnregisterJobs(jobIds []models.JobId) error
	UpdateJobs(jobs []models.Job) error

	LaunchJob(jobId models.JobId) error
	TerminateJob(jobId models.JobId) error

	FindJob(jobId models.JobId) (models.Job, error)
	FindAllJobs() ([]models.Job, error)
	FindJobsByStatus(status models.JobStatus) ([]models.Job, error)

	HandleProcessEvent(event models.ProcessEvent) error
}
