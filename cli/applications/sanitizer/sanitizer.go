package sanitizer

import (
	"strings"

	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/core/error_msg"
	"golang.org/x/tools/go/analysis/passes/unreachable"
)

type SimpleSanitizer struct {
	SupervisorTranslator interfaces.ISupervisorTranslator
	SupervisorAdapter    interfaces.ISupervisorAdapter
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

func (s *SimpleSanitizer) Run(parsedComand interfaces.IParsedCommand) (interfaces.ISanitizedCommand, interfaces.IResponse) {
	command := ValidateCommand(parsedComand.Command())
	resp_builder := domain.NewResponseBuilder()

	if command == CmdInvalid {
		resp_builder.Error(parsedComand.Command() + " " + error_msg.BAD_COMMAND)
		return nil, resp_builder.Build()
	}

	if command == CmdHelp {
		resp_builder.Info("Available commands : " + strings.Join(CommandSymbols[1:], ", "))
		return nil, resp_builder.Build()
	}

	// Get la liste des jobs et s'assurer que chaque nom de jobs existe

	// Get job status et faire en sorte que la commande est executable

	statuses, error := s.SupervisorAdapter.GetJobStatuses(parsedComand.JobNames())
	if error != nil {
		resp_builder.Error(error_msg.INTERNAL_ERROR + " " + error.Error())
		return nil, resp_builder.Build()
	}

	possible_statuses := possibleStatuses(command)

	return &SanitizedCommand{
		command: command,
		jobs:    []interfaces.IJob{},
	}, nil
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
	_ domain.Status = iota
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

func possibleStatuses(code domain.CommandCode) []domain.Status {
	ret_val := []domain.Status{}
	switch code {
	case CmdExit:
		ret_val = []domain.Status{}
	case CmdPid:
		ret_val = []domain.Status{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	case CmdHelp:
		ret_val = []domain.Status{}
	case CmdUpdate:
		ret_val = []domain.Status{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	case CmdStart:
		ret_val = []domain.Status{Stoped, Stoping, Running, Exited, Fatal}
	case CmdStop:
		ret_val = []domain.Status{Stoped, Starting, Running}
	case CmdRestart:
		ret_val = []domain.Status{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	case CmdStatus:
		ret_val = []domain.Status{Stoped, Starting, Stoping, Backoff, Running, Exited, Fatal}
	}
	return ret_val
}
