package infrastructure

import (
	"sync"

	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/services"
)

type ProcessWatcher struct {
	processManager *OSProcessManager
	jobService     services.IJobService
	wg             sync.WaitGroup
}

func NewProcessWatcher(processManager services.IProcessManager, jobService services.IJobService) services.IWatcherService {
	return &ProcessWatcher{
		processManager: processManager.(*OSProcessManager),
		jobService:     jobService,
	}
}

func (pw *ProcessWatcher) Start() error                          { return nil }
func (pw *ProcessWatcher) Stop() error                           { pw.wg.Wait(); return nil }
func (pw *ProcessWatcher) UnregisterJob(models.JobId) error      { return nil }
func (pw *ProcessWatcher) CheckHealth() *models.JobsHealthResult { return nil }

func (pw *ProcessWatcher) RegisterJob(job *models.Job) error {
	if job.IsAlive() && job.Pid > 0 {
		pw.wg.Add(1)
		go pw.monitorJob(job.Id)
	}
	return nil
}

func (pw *ProcessWatcher) monitorJob(jobId models.JobId) {
	defer pw.wg.Done()
	exitCode, err, completed := pw.processManager.WaitForExit(jobId)

	var event models.ProcessEvent
	if err != nil && !completed {
		event = models.NewProcessFailedEvent(jobId, 0, err)
	} else {
		event = models.NewProcessExitedEvent(jobId, 0, exitCode)
	}

	pw.jobService.HandleProcessEvent(event)
}
