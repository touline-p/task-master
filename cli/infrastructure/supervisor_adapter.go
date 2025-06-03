package infrastructure

import (
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	supModels "github.com/touline-p/task-master/supervisor/domain/models"
)

// ACL
type SupervisorTranslator struct {}

func (st *SupervisorTranslator)Translate(strings []string) []string {return strings}

type SupervisorAdapter struct {
	commandHandler cqrs.ICommandHandler
	queryHandler   cqrs.IQueryHandler
}

func NewSupervisorAdapter(commandHandler cqrs.ICommandHandler, queryHandler cqrs.IQueryHandler) *SupervisorAdapter {
	return &SupervisorAdapter{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

func (a *SupervisorAdapter) StartJobs(jobIds []string) error {
	for _, jobId := range(jobIds) {
		a.StartJob(jobId)
	}
	return nil
}

func (a *SupervisorAdapter) StopJobs(jobIds []string) error {
	for _, jobId := range(jobIds) {
		a.StopJob(jobId)
	}
	return nil
}
func (a *SupervisorAdapter) StartJob(jobId string) error {
	cmd := &cqrs.StartJobCommand{JobId: supModels.JobId(jobId)}
	return a.commandHandler.HandleStartJob(cmd)
}

func (a *SupervisorAdapter) StopJob(jobId string) error {
	cmd := &cqrs.StopJobCommand{JobId: supModels.JobId(jobId)}
	return a.commandHandler.HandleStopJob(cmd)
}

func (a *SupervisorAdapter) RestartJob(jobId string) error {
	cmd := &cqrs.RestartJobCommand{JobId: supModels.JobId(jobId)}
	return a.commandHandler.HandleRestartJob(cmd)
}

func (a *SupervisorAdapter) RestartJobs(jobIds[] string) error {
	for _, jobId := range(jobIds) {
		a.RestartJob(jobId)
	}
	return nil
}

// Translate supervisor -> cli
func (a *SupervisorAdapter) GetJobStatuses(jobIds[]string) (map[string]interfaces.JobStatus, error) {
	query := &cqrs.GetJobStatusesQuery{}
	result, err := a.queryHandler.HandleGetJobStatuses(query)
	if err != nil {
		return nil, err
	}

	statuses := make(map[string]interfaces.JobStatus)

	for id, state := range result {
		cliStatus := mapSupervisorStatus(state.Status)
		statuses[string(id)] = cliStatus
	}

	return statuses, nil
}

func mapSupervisorStatus(status supModels.JobStatus) interfaces.JobStatus {
	switch status {
	case supModels.StatusStopped:
		return interfaces.StatusStopped
	case supModels.StatusStarting:
		return interfaces.StatusStarting
	case supModels.StatusRunning:
		return interfaces.StatusRunning
	case supModels.StatusBackoff:
		return interfaces.StatusBackoff
	case supModels.StatusStopping:
		return interfaces.StatusStopping
	case supModels.StatusExited:
		return interfaces.StatusExited
	case supModels.StatusFatal:
		return interfaces.StatusFatal
	default:
		return interfaces.StatusStopped
	}
}
