package services

import (
	"context"
	"fmt"

	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/services"
)

type JobService struct {
	processManager services.IProcessManager
}

func NewJobService(processManager services.IProcessManager) *JobService {
	return &JobService{
		processManager: processManager,
	}
}

func (s *JobService) StartJob(ctx context.Context, job *models.Job) error {
	if err := job.Start(); err != nil {
		return err
	}

	pid, err := s.processManager.Launch(ctx, job)
	if err != nil {
		job.MarkAsFatal(fmt.Sprintf("Failed to start process: %v", err))
		return err
	}

	if err := job.MarkAsRunning(pid); err != nil {
		return err
	}

	return nil
}

func (s *JobService) StopJob(ctx context.Context, job *models.Job) error {
	if err := job.Stop(); err != nil {
		return err
	}

	if err := job.MarkAsStopped(); err != nil {
		return err
	}

	return nil
}
