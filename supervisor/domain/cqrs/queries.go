package cqrs

import (
	"github.com/touline-p/task-master/supervisor/domain/models"
)

type Query interface {
	Type() string
}

type GetJobStatusesQuery struct{}

func (q *GetJobStatusesQuery) Type() string {
	return "GET_JOB_STATUSES"
}

type GetJobByIdQuery struct {
	JobId models.JobId
}

func (q *GetJobByIdQuery) Type() string {
	return "GET_JOB_BY_ID"
}

type GetJobsByStatusQuery struct {
	Status models.JobStatus
}

func (q *GetJobsByStatusQuery) Type() string {
	return "GET_JOBS_BY_STATUS"
}

type CheckJobHealthQuery struct{}

func (q *CheckJobHealthQuery) Type() string {
	return "CHECK_JOB_HEALTH"
}

type IQueryHandler interface {
	HandleGetJobStatuses(query *GetJobStatusesQuery) (map[models.JobId]models.JobState, error)
	HandleGetJobById(query *GetJobByIdQuery) (models.Job, error)
	HandleGetJobsByStatus(query *GetJobsByStatusQuery) ([]models.Job, error)
	HandleCheckJobHealth(query *CheckJobHealthQuery) (map[models.JobId]bool, error)
}
