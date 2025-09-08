package cqrs

import (
	"context"
	"fmt"

	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/services"
)

type JobCommandHandler struct {
	jobService       services.IJobService
	schedulerService services.ISchedulerService
}

func NewJobCommandHandler(jobService services.IJobService) cqrs.ICommandHandler {
	return &JobCommandHandler{
		jobService:       jobService,
		schedulerService: nil,
	}
}

func (h *JobCommandHandler) SetSchedulerService(schedulerService services.ISchedulerService) {
	h.schedulerService = schedulerService
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

func (h *JobCommandHandler) HandleStopSupervisor(cmd *cqrs.StopSupervisorCommand) error {
	if h.schedulerService == nil {
		return fmt.Errorf("scheduler service not initialized")
	}
	return h.schedulerService.Shutdown()
}
