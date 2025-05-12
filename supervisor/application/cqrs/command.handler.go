package cqrs

import (
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
)

type JobCommandHandler struct {
	repository repositories.IJobRepository
}

func NewJobCommandHandler(repository repositories.IJobRepository) cqrs.ICommandHandler {
	return &JobCommandHandler{
		repository: repository,
	}
}

func (h *JobCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.(type) {
	case *cqrs.StartJobCommand:
		return h.handleStartJob(cmd)
	case *cqrs.StopJobCommand:
		return h.handleStopJob(cmd)
	case *cqrs.RestartJobCommand:
		return h.handleRestartJob(cmd)
	case *cqrs.ProcessEventCommand:
		return h.handleProcessEvent(cmd)
	default:
		return nil
	}
}

func (h *JobCommandHandler) handleStartJob(cmd *cqrs.StartJobCommand) error {
	job, err := h.repository.FindById(cmd.JobId)
	if err != nil {
		return err
	}

	err = job.Start()
	if err != nil {
		return err
	}

	return h.repository.Save(&job)
}

func (h *JobCommandHandler) handleStopJob(cmd *cqrs.StopJobCommand) error {
	job, err := h.repository.FindById(cmd.JobId)
	if err != nil {
		return err
	}

	err = job.Stop()
	if err != nil {
		return err
	}

	return h.repository.Save(&job)
}

func (h *JobCommandHandler) handleRestartJob(cmd *cqrs.RestartJobCommand) error {
	job, err := h.repository.FindById(cmd.JobId)
	if err != nil {
		return err
	}

	err = job.Stop()
	if err != nil {
		return err
	}

	err = h.repository.Save(&job)
	if err != nil {
		return err
	}

	err = job.Start()
	if err != nil {
		return err
	}

	return h.repository.Save(&job)
}

func (h *JobCommandHandler) handleProcessEvent(cmd *cqrs.ProcessEventCommand) error {
	job, err := h.repository.FindById(cmd.Event.JobId)
	if err != nil {
		return err
	}

	switch cmd.Event.EventType {
	case models.ProcessStarted:
		err = job.MarkAsRunning(0) // TODO : pid
	case models.ProcessExited:
		err = job.MarkAsExited(cmd.Event.ExitCode)
		if err == nil && job.ShouldRestart(cmd.Event.ExitCode) {
			if job.HasExceededRetries() {
				err = job.MarkAsFatal("Exceeded maximum retry attempts")
			} else {
				err = job.BackOff()
			}
		}
	case models.ProcessFailed:
		if cmd.Event.Error != nil {
			err = job.MarkAsFatal(cmd.Event.Error.Error())
		} else {
			err = job.MarkAsFatal("Process failed")
		}
	}

	if err != nil {
		return err
	}

	return h.repository.Save(&job)
}
