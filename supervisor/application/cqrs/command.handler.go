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

func (h *JobCommandHandler) HandleStartJob(cmd *cqrs.StartJobCommand) error {
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

func (h *JobCommandHandler) HandleStopJob(cmd *cqrs.StopJobCommand) error {
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

func (h *JobCommandHandler) HandleRestartJob(cmd *cqrs.RestartJobCommand) error {
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

func (h *JobCommandHandler) HandleProcessEvent(cmd *cqrs.ProcessEventCommand) error {
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
