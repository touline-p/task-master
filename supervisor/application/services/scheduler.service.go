package services

import (
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
	"github.com/touline-p/task-master/supervisor/domain/services"
)

type SchedulerService struct {
	repository     repositories.IJobRepository
	commandHandler cqrs.ICommandHandler
	queryHandler   cqrs.IQueryHandler
}

func NewSchedulerService(repository repositories.IJobRepository, commandHandler cqrs.ICommandHandler, queryHandler cqrs.IQueryHandler) services.ISchedulerService {
	return &SchedulerService{
		repository:     repository,
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

func (ss *SchedulerService) RegisterJobs(jobs []models.Job) error {
	errors := make([]error, 0)
	for _, j := range jobs {
		err := ss.repository.Save(&j)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return ConcatenateErrors(errors)
}
