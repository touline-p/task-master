package cqrs

import (
	"context"

	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/services"
)

type JobCommandHandler struct {
	jobService services.IJobService
}

func NewJobCommandHandler(jobService services.IJobService) cqrs.ICommandHandler {
	return &JobCommandHandler{
		jobService: jobService,
	}
}

func (h *JobCommandHandler) HandleStartJob(cmd *cqrs.StartJobCommand) error {
	ctx := context.Background()
	return h.jobService.StartJob(ctx, cmd.JobId)
}

func (h *JobCommandHandler) HandleStopJob(cmd *cqrs.StopJobCommand) error {
	return h.jobService.StopJob(cmd.JobId)
}

func (h *JobCommandHandler) HandleRestartJob(cmd *cqrs.RestartJobCommand) error {
	ctx := context.Background()

	err := h.jobService.StopJob(cmd.JobId)
	if err != nil {
		return err
	}

	return h.jobService.StartJob(ctx, cmd.JobId)
}

func (h *JobCommandHandler) HandleProcessEvent(cmd *cqrs.ProcessEventCommand) error {
	return h.jobService.HandleProcessEvent(cmd.Event)
}
