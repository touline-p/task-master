package launcher

import (
	"github.com/touline-p/task-master/cli/applications/sanitizer"
	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleLauncher struct {
	SupervisorAdapter interfaces.ISupervisorAdapter
}

func (l *SimpleLauncher) Run(cmd interfaces.ISanitizedCommand) interfaces.IResponse {
	builder := domain.NewResponseBuilder()

	if cmd == nil {
		builder.Error(error_msg.BAD_COMMAND)
		return builder.Build()
	}

	switch cmd.Code() {
	case sanitizer.CmdStart:
		for _, jobId := range cmd.JobIds() {
			jobIdStr := string(jobId.(domain.JobId))
			err := l.SupervisorAdapter.StartJob(jobIdStr)
			if err != nil {
				builder.Error(err.Error())
			}
		}
	case sanitizer.CmdStop:
		for _, jobId := range cmd.JobIds() {
			jobIdStr := string(jobId.(domain.JobId))
			err := l.SupervisorAdapter.StopJob(jobIdStr)
			if err != nil {
				builder.Error(err.Error())
			}
		}
	case sanitizer.CmdRestart:
		for _, jobId := range cmd.JobIds() {
			jobIdStr := string(jobId.(domain.JobId))
			err := l.SupervisorAdapter.RestartJob(jobIdStr)
			if err != nil {
				builder.Error(err.Error())
			}
		}
	case sanitizer.CmdStatus:
		statuses, err := l.SupervisorAdapter.GetJobStatuses()
		if err != nil {
			builder.Error(err.Error())
		} else {
			// Format the statuses for output
			for id, status := range statuses {
				builder.Warning(id + ": " + statusToString(status))
			}
		}
	default:
		builder.Error(error_msg.BAD_COMMAND)
	}

	return builder.Build()
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
