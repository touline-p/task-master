package services

import (
	"context"
	"fmt"

	"github.com/touline-p/task-master/supervisor/application/ports"
	"github.com/touline-p/task-master/supervisor/domain/models"
)

type JobService struct {
	processManager ports.ProcessManager
}

func NewJobService(processManager ports.ProcessManager) *JobService {
	return &JobService{
		processManager: processManager,
	}
}

func (s *JobService) StartJob(ctx context.Context, job *models.Job) error {
	if err := job.Start(); err != nil {
		return err
	}

	pid, err := s.processManager.SpawnProcess(ctx, job)
	if err != nil {
		job.MarkAsFatal(fmt.Sprintf("Failed to start process: %v", err))
		return err
	}

	if err := job.MarkAsRunning(pid); err != nil {
		return err
	}

	return nil
}
