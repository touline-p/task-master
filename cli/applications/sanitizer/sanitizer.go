package sanitizer

import (
	"strings"

	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleSanitizer struct{}

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

func (self *SimpleSanitizer) Run(parsedComand interfaces.IParsedCommand) (interfaces.ISanitizedCommand, interfaces.IResponse) {
	command := ValidateCommand(parsedComand.Command())
	resp_builder := domain.NewResponseBuilder()

	if command == CmdInvalid {
		resp_builder.Error("command: " + parsedComand.Command() + " " + error_msg.BAD_COMMAND)
		return nil, resp_builder.Build()
	}

	if command == CmdHelp {
		resp_builder.Info("Available commands : " + strings.Join(CommandSymbols[1:], ", "))
		return nil, resp_builder.Build()
	}
	// Get job status et faire en sorte que la commande est executable

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
