package interfaces

import (
	models "github.com/touline-p/task-master/supervisor/domain/models"
)

type ISchedulerService interface {
	RunningJobs() []models.Job
	BackedoffJobs() []models.Job
	addJob(models.Job)
}
