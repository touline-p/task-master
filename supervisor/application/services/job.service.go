package services

import (
	"context"
	"fmt"

	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
	"github.com/touline-p/task-master/supervisor/domain/services"
)

type JobService struct {
	processManager services.IProcessManager
	repository     repositories.IJobRepository
}

func NewJobService(processManager services.IProcessManager, repository repositories.IJobRepository) *JobService {
	return &JobService{
		processManager: processManager,
		repository:     repository,
	}
}

func (s *JobService) StartJob(ctx context.Context, jobId models.JobId) error {
	job, err := s.repository.FindById(jobId)
	if err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	return s.withSave(&job, func() error {
		if err := job.Start(); err != nil {
			return err
		}

		pid, err := s.processManager.Launch(ctx, &job)
		if err != nil {
			job.MarkAsFatal(fmt.Sprintf("Failed to start process: %v", err))
			return err
		}

		return job.MarkAsRunning(pid)
	})
}

func (s *JobService) StopJob(jobId models.JobId) error {
	job, err := s.repository.FindById(jobId)
	if err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	return s.withSave(&job, func() error {
		if err := job.Stop(); err != nil {
			return err
		}

		if job.Pid > 0 {
			if err := s.processManager.Terminate(&job); err != nil {
				if killErr := s.processManager.Kill(&job); killErr != nil {
					return fmt.Errorf("failed to stop process: terminate failed (%v), kill failed (%v)", err, killErr)
				}
			}
		}

		return job.MarkAsStopped()
	})
}

func (s *JobService) HandleProcessEvent(event models.ProcessEvent) error {
	job, err := s.repository.FindById(event.JobId)
	if err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	return s.withSave(&job, func() error {
		switch event.EventType {
		case models.ProcessStarted:
			err = job.MarkAsRunning(0) // TODO : pid
		case models.ProcessExited:
			err = job.MarkAsExited(event.ExitCode)
			if err == nil && job.ShouldRestart(event.ExitCode) {
				if job.HasExceededRetries() {
					err = job.MarkAsFatal("Exceeded maximum retry attempts")
				} else {
					err = job.BackOff()
				}
			}
		case models.ProcessFailed:
			if event.Error != nil {
				err = job.MarkAsFatal(event.Error.Error())
			} else {
				err = job.MarkAsFatal("Process failed")
			}
		}

		if err != nil {
			return err
		}
		return nil
	})
}

func (s *JobService) withSave(job *models.Job, operation func() error) error {
	if err := operation(); err != nil {
		return err
	}
	return s.repository.Save(job)
}
