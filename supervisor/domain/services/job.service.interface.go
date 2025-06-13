package services

import (
	"context"

	"github.com/touline-p/task-master/supervisor/domain/models"
)

type IJobService interface {
	StartJob(ctx context.Context, jobId models.JobId) error
	StopJob(jobId models.JobId) error
	HandleProcessEvent(event models.ProcessEvent) error
}
