package repositories

import (
	models "github.com/touline-p/task-master/supervisor/domain/models"
)

type ISchedulerRepository interface {
	create() (*models.Job, error)
	getById(id models.JobId) (*models.Job, error)
	getAll() ([]*models.Job, error)
	update() (*models.Job, error)
	delete(*models.Job) error
}
