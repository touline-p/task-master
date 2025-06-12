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
	watcher        svcInterfaces.IWatcherService
	jobService     svcInterfaces.IJobService
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

func (c *Controller) JobService() svcInterfaces.IJobService {
	return c.jobService
}

func (c *Controller) Watcher() svcInterfaces.IWatcherService {
	return c.watcher
}

func GetSupervisorController() *Controller {
	repository := infrastructure.GetJobRepository()
	queryHandler := appCqrs.NewJobQueryHandler(repository)
	processManager := infrastructure.NewOSProcessManager()
	jobService := services.NewJobService(processManager, repository)
	commandHandler := appCqrs.NewJobCommandHandler(jobService)
	scheduler := services.NewSchedulerService(repository, commandHandler, queryHandler)
	watcher := infrastructure.NewProcessWatcher(processManager, jobService)

	return &Controller{
		repository:     repository,
		queryHandler:   queryHandler,
		processManager: processManager,
		watcher:        watcher,
		jobService:     jobService,
		commandHandler: commandHandler,
		scheduler:      scheduler,
	}
}
