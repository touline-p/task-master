package cqrs

import (
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
)

type JobQueryHandler struct {
	repository repositories.IJobRepository
}

func NewJobQueryHandler(repository repositories.IJobRepository) cqrs.IQueryHandler {
	return &JobQueryHandler{
		repository: repository,
	}
}

func (h *JobQueryHandler) HandleGetJobStatuses(q *cqrs.GetJobStatusesQuery) (map[models.JobId]models.JobState, error) {
	jobs, err := h.repository.FindAll()
	if err != nil {
		return nil, err
	}

	statuses := make(map[models.JobId]models.JobState)
	for _, job := range jobs {
		statuses[job.Id] = *job.GetState()
	}

	return statuses, nil
}

func (h *JobQueryHandler) HandleGetJobById(q *cqrs.GetJobByIdQuery) (models.Job, error) {
	return h.repository.FindById(q.JobId)
}

func (h *JobQueryHandler) HandleGetJobsByStatus(q *cqrs.GetJobsByStatusQuery) ([]models.Job, error) {
	return h.repository.FindByStatus(q.Status)
}

func (h *JobQueryHandler) HandleCheckJobHealth(q *cqrs.CheckJobHealthQuery) (map[models.JobId]bool, error) {
	jobs, err := h.repository.FindAll()
	if err != nil {
		return nil, err
	}

	healthReport := make(map[models.JobId]bool)
	for _, job := range jobs {
		// A simple health check: a job is healthy if it's in the running state
		healthReport[job.Id] = job.GetState().Status == models.StatusRunning
	}

	return healthReport, nil
}
