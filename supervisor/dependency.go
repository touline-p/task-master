package supervisor

import (
	appCqrs "github.com/touline-p/task-master/supervisor/application/cqrs"
	"github.com/touline-p/task-master/supervisor/application/services"
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
	svcInterfaces "github.com/touline-p/task-master/supervisor/domain/services"
	"github.com/touline-p/task-master/supervisor/infrastructure"
)

type Controller struct {
	repository     repositories.IJobRepository
	scheduler      svcInterfaces.ISchedulerService
	commandHandler cqrs.ICommandHandler
	queryHandler   cqrs.IQueryHandler
	processManager svcInterfaces.IProcessManager
	jobService     *services.JobService
}

func (c *Controller) Repository() repositories.IJobRepository {
	return c.repository
}

func (c *Controller) Scheduler() svcInterfaces.ISchedulerService {
	return c.scheduler
}

func (c *Controller) CommandHandler() cqrs.ICommandHandler {
	return c.commandHandler
}

func (c *Controller) QueryHandler() cqrs.IQueryHandler {
	return c.queryHandler
}

func (c *Controller) ProcessManager() svcInterfaces.IProcessManager {
	return c.processManager
}

func (c *Controller) JobService() *services.JobService {
	return c.jobService
}

func GetSupervisorController() *Controller {
	repository := infrastructure.NewJobRepository()
	commandHandler := appCqrs.NewJobCommandHandler(repository)
	queryHandler := appCqrs.NewJobQueryHandler(repository)
	processManager := infrastructure.NewOSProcessManager()
	jobService := services.NewJobService(processManager, repository)

	return &Controller{
		repository:     repository,
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
		scheduler:      services.NewSchedulerService(repository, commandHandler, queryHandler),
		processManager: processManager,
		jobService:     jobService,
	}
}
