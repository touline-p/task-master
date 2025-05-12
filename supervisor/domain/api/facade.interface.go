package api

import (
	"github.com/touline-p/task-master/supervisor/domain/models"
)

type ISupervisorFacade interface {
	GetJobStatuses() ([]*models.Job, error)
	GetJobById(jobId models.JobId) (*models.Job, error)
	GetJobsByStatus(status models.JobStatus) ([]*models.Job, error)

	StartJob(jobId models.JobId) error
	StopJob(jobId models.JobId) error
	RestartJob(jobId models.JobId) error

	ReloadConfig(configPath string) error
	CheckHealth() *models.JobsHealthResult
}
