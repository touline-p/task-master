package launcher

import (
	"github.com/touline-p/task-master/cli/applications/sanitizer"
	"github.com/touline-p/task-master/cli/domain"
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

func (l *SimpleLauncher) Run(cmd interfaces.ISanitizedCommand) interfaces.IResponse {
	resp_bldr := domain.NewResponseBuilder()
	translator := l.SvTranslator()
	adapter := l.SvAdapter()

	if cmd == nil {
		resp_bldr.Error(error_msg.BAD_COMMAND)
		return resp_bldr.Build()
	}

	stringifiedJobs := []string{}
	for _, jobId := range cmd.JobIds() {
		stringifiedJobs = append(stringifiedJobs, jobId.ToString())
	}

	translated_job := translator.Translate(stringifiedJobs)

	switch cmd.Code() {
	case sanitizer.CmdStart:
		resp_bldr.HandleCmd(adapter.StartJobs(translated_job))
	case sanitizer.CmdStop:
		resp_bldr.HandleCmd(adapter.StopJobs(translated_job))
	case sanitizer.CmdRestart:
		resp_bldr.HandleCmd(adapter.RestartJobs(translated_job))
	case sanitizer.CmdStatus:
		query, err := adapter.GetJobStatuses(translated_job)
		for key, value := range query {
			resp_bldr.Info(key + statusToString(value))
		}
		resp_bldr.Error(err.Error())
	default:
		resp_bldr.Error(error_msg.BAD_COMMAND)
	}
	return resp_bldr.Build()
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
