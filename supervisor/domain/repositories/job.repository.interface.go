package repositories

import (
	models "github.com/touline-p/task-master/supervisor/domain/models"
)

type IJobRepository interface {
	Save(job *models.Job) error
	FindById(id models.JobId) (*models.Job, error)
	FindAll() ([]*models.Job, error)
	FindByStatus(status models.JobStatus) ([]*models.Job, error)
	Delete(id models.JobId) error
	Exists(id models.JobId) bool
}
