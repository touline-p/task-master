package sanitizer

import (
	"fmt"
	"slices"
	"strings"

	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleSanitizer struct {
	SupervisorTranslator interfaces.ISupervisorTranslator
	SupervisorAdapter    interfaces.ISupervisorAdapter
}

type Job struct {
	id string
}

func (j *Job) Id() string {
	return j.id
}

type SanitizedCommand struct {
	command domain.CommandCode
	jobs    []interfaces.IJob
}

func (sc *SanitizedCommand) Code() interfaces.ICommandCode {
	return sc.command
}

func (sc *SanitizedCommand) JobIds() []interfaces.IJob {
	return sc.jobs
}

func (s *SimpleSanitizer) Run(parsedComand interfaces.IParsedCommand, bldr interfaces.IResponseBuilder) (interfaces.ISanitizedCommand, interfaces.IResponseBuilder) {
	command := ValidateCommand(parsedComand.Command())

	if command == CmdInvalid {
		bldr.Error(parsedComand.Command() + " " + error_msg.BAD_COMMAND)
		return nil, bldr
	}

	if command == CmdHelp {
		bldr.Info("Available commands : " + strings.Join(CommandSymbols[1:], ", "))
		return nil, bldr
	}

	statuses, err := s.SupervisorAdapter.GetJobStatuses([]string{})
	if err != nil {
		bldr.Error(err.Error())
		return nil, bldr
	}
	filteredJobStatus, bldr := filterUnknownJobs(parsedComand.JobNames(), statuses, bldr)
	validated_jobs, bldr := validatePossibleJobs(command, filteredJobStatus, bldr)

	return &SanitizedCommand{
		command: command,
		jobs:    validated_jobs,
	}, bldr
}

func filterUnknownJobs(jobNames []string, statuses map[string]interfaces.JobStatus, resp_builder interfaces.IResponseBuilder) (map[string]interfaces.JobStatus, interfaces.IResponseBuilder) {
	filtered_statuses := make(map[string]interfaces.JobStatus)
	for _, n := range jobNames {
		status, exist := statuses[n]
		if exist == true {
			filtered_statuses[n] = status
		} else {
			resp_builder.Warning(n + ": " + error_msg.UNKNOWN_JOB)
		}
	}
	return filtered_statuses, resp_builder
}

func validatePossibleJobs(command domain.CommandCode, statuses map[string]interfaces.JobStatus, resp_builder interfaces.IResponseBuilder) ([]interfaces.IJob, interfaces.IResponseBuilder) {
	id_array := []interfaces.IJob{}
	possible_statuses := possibleStatuses(command)
	for id, status := range statuses {
		if slices.Contains(possible_statuses, status) {
			id_array = append(id_array, &Job{id: id})
		} else {
			resp_builder.Warning(fmt.Sprintf("%s is %s. It can't be %s.", id, status, command))
		}
	}
	return id_array, resp_builder
}

const (
	CmdInvalid domain.CommandCode = iota
	CmdExit
	CmdPid
	CmdHelp
	CmdUpdate
	CmdStart
	CmdStop
	CmdRestart
	CmdStatus
	CmdNumber
)

const (
	_ interfaces.JobStatus = iota
	Stoped
	Starting
	Stoping
	Backoff
	Running
	Exited
	Fatal
)

var CommandSymbols = [CmdNumber]string{
	"",
	"exit",
	"pid",
	"help",
	"update",
	"start",
	"stop",
	"restart",
	"status",
}

func ValidateCommand(command string) domain.CommandCode {
	for commandCode, symbol := range CommandSymbols {
		if symbol == command {
			return domain.CommandCode(commandCode)
		}
	}
	return CmdInvalid
}

func possibleStatuses(code domain.CommandCode) []interfaces.JobStatus {
	switch code {
	case CmdExit:
		return []interfaces.JobStatus{}
	case CmdPid:
		return []interfaces.JobStatus{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	case CmdHelp:
		return []interfaces.JobStatus{}
	case CmdUpdate:
		return []interfaces.JobStatus{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	case CmdStart:
		return []interfaces.JobStatus{Stoped, Stoping, Running, Exited, Fatal}
	case CmdStop:
		return []interfaces.JobStatus{Stoped, Starting, Running}
	case CmdRestart:
		return []interfaces.JobStatus{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	case CmdStatus:
		return []interfaces.JobStatus{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	}
	return []interfaces.JobStatus{}
}
