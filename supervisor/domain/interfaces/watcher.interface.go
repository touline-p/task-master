package interfaces

import (
	models "github.com/touline-p/task-master/supervisor/domain/models"
)

type IWatcherService interface {
	checkJobsHealth() models.JobsHealthResult
}
