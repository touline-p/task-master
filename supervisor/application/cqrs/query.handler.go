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

func (h *JobQueryHandler) Handle(query cqrs.Query) (any, error) {
	switch q := query.(type) {
	case *cqrs.GetJobStatusesQuery:
		return h.handleGetJobStatuses(q)
	case *cqrs.GetJobByIdQuery:
		return h.handleGetJobById(q)
	case *cqrs.GetJobsByStatusQuery:
		return h.handleGetJobsByStatus(q)
	case *cqrs.CheckJobHealthQuery:
		return h.handleCheckJobHealth(q)
	default:
		return nil, nil
	}
}

func (h *JobQueryHandler) handleGetJobStatuses(q *cqrs.GetJobStatusesQuery) (any, error) {
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

func (h *JobQueryHandler) handleGetJobById(q *cqrs.GetJobByIdQuery) (any, error) {
	return h.repository.FindById(q.JobId)
}

func (h *JobQueryHandler) handleGetJobsByStatus(q *cqrs.GetJobsByStatusQuery) (any, error) {
	return h.repository.FindByStatus(q.Status)
}

func (h *JobQueryHandler) handleCheckJobHealth(q *cqrs.CheckJobHealthQuery) (any, error) {
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
