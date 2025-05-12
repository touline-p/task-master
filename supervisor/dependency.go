package supervisor

import (
	appCqrs "github.com/touline-p/task-master/supervisor/application/cqrs"
	appServices "github.com/touline-p/task-master/supervisor/application/services"
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
	"github.com/touline-p/task-master/supervisor/domain/services"
	"github.com/touline-p/task-master/supervisor/infrastructure"
)

type Controller struct {
	repository     repositories.IJobRepository
	scheduler      services.ISchedulerService
	commandHandler cqrs.ICommandHandler
	queryHandler   cqrs.IQueryHandler
}

func (c *Controller) Repository() repositories.IJobRepository {
	return c.repository
}

func (c *Controller) Scheduler() services.ISchedulerService {
	return c.scheduler
}

func (c *Controller) CommandHandler() cqrs.ICommandHandler {
	return c.commandHandler
}

func (c *Controller) QueryHandler() cqrs.IQueryHandler {
	return c.queryHandler
}

func GetSupervisorController() *Controller {
	repository := infrastructure.NewJobRepository()
	commandHandler := appCqrs.NewJobCommandHandler(repository)
	queryHandler := appCqrs.NewJobQueryHandler(repository)

	return &Controller{
		repository:     repository,
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
		scheduler:      appServices.NewSchedulerService(repository, commandHandler, queryHandler),
	}
}
