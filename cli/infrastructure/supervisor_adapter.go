package infrastructure

import (
	"fmt"

	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	supModels "github.com/touline-p/task-master/supervisor/domain/models"
)

// ACL
type SupervisorTranslator struct{}

func (st *SupervisorTranslator) Translate(strings []string) []string { return strings }

type CliJob struct {
	id string
}

func (j *CliJob) Id() string {
	return j.id
}

func (st *SupervisorTranslator) StringToJob(strings []string) []interfaces.IJob {
	jobs := make([]interfaces.IJob, len(strings))
	for i, s := range strings {
		jobs[i] = &CliJob{id: s}
	}
	return jobs
}
func (st *SupervisorTranslator) JobToString(jobs []interfaces.IJob) []string {
	strings := make([]string, len(jobs))
	for i, j := range jobs {
		strings[i] = j.Id()
	}
	return strings
}

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
	for _, jobId := range jobIds {
		a.StartJob(jobId)
	}
	return nil
}

func (a *SupervisorAdapter) StopJobs(jobIds []string) error {
	for _, jobId := range jobIds {
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

func (a *SupervisorAdapter) RestartJobs(jobIds []string) error {
	for _, jobId := range jobIds {
		a.RestartJob(jobId)
	}
	return nil
}

// Translate supervisor -> cli
func (a *SupervisorAdapter) GetJobStatuses(jobIds []string) (map[string]interfaces.JobStatus, error) {
	query := &cqrs.GetJobStatusesQuery{}
	result, err := a.queryHandler.HandleGetJobStatuses(query)
	if err != nil {
		println("an error occured, fuck.")
		return nil, err
	}

	statuses := make(map[string]interfaces.JobStatus)

	for id, state := range result {
		cliStatus := mapSupervisorStatus(state.Status)
		statuses[string(id)] = cliStatus
	}
	println(len(statuses))
	for key, value := range statuses {
		fmt.Println("%s, %d", key, value)
	}

	if len(jobIds) > 0 {
		filterStatusesMap(statuses, jobIds)
	}
	return statuses, nil
}

func filterStatusesMap(m map[string]interfaces.JobStatus, jobIds []string) {
	keepSet := make(map[string]bool)
	for _, key := range jobIds {
		keepSet[key] = true
	}

	for key := range m {
		if !keepSet[key] {
			delete(m, key)
		}
	}
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
