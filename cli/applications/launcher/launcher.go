package launcher

import (
	"github.com/touline-p/task-master/cli/applications/sanitizer"
	"github.com/touline-p/task-master/cli/domain/interfaces"

	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleLauncher struct {
	SupervisorTranslator interfaces.ISupervisorTranslator
	SupervisorAdapter    interfaces.ISupervisorAdapter
}

func (sl *SimpleLauncher) SvTranslator() interfaces.ISupervisorTranslator {
	return sl.SupervisorTranslator
}
func (sl *SimpleLauncher) SvAdapter() interfaces.ISupervisorAdapter { return sl.SupervisorAdapter }

func (l *SimpleLauncher) Run(cmd interfaces.ISanitizedCommand, bldr interfaces.IResponseBuilder) interfaces.IResponseBuilder {
	translator := l.SvTranslator()
	adapter := l.SvAdapter()

	stringifiedJobs := []string{}
	for _, jobId := range translator.JobToString(cmd.JobIds()) {
		stringifiedJobs = append(stringifiedJobs, jobId)
	}

	translated_job := translator.Translate(stringifiedJobs)

	var err error
	switch cmd.Code() {
	case sanitizer.CmdStart:
		err = adapter.StartJobs(translated_job)
	case sanitizer.CmdStop:
		err = adapter.StopJobs(translated_job)
	case sanitizer.CmdRestart:
		err = adapter.RestartJobs(translated_job)
	case sanitizer.CmdStatus:
		var query map[string]interfaces.JobStatus
		query, err = adapter.GetJobStatuses(translated_job)
		for key, value := range query {
			bldr.Info(key + statusToString(value))
		}
	default:
		bldr.Error(error_msg.BAD_COMMAND)
	}
	if err != nil {
		bldr.Error(err.Error())
	}
	return bldr
}

func statusToString(status interfaces.JobStatus) string {
	switch status {
	case interfaces.StatusStopped:
		return "stopped"
	case interfaces.StatusStarting:
		return "starting"
	case interfaces.StatusRunning:
		return "running"
	case interfaces.StatusStopping:
		return "stopping"
	case interfaces.StatusBackoff:
		return "backoff"
	case interfaces.StatusExited:
		return "exited"
	case interfaces.StatusFatal:
		return "fatal"
	default:
		return "unknown"
	}
}
