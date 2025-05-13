package sanitizer

import (
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleSanitizer struct{}

type SanitizedCommand struct {
	command interfaces.ICommandCode
	jobs    []interfaces.IJob
}

func (self *SimpleSanitizer) Run(parsedComand interfaces.IParsedCommand) (interfaces.ISanitizedCommand, interfaces.IResponse) {
	command := ValidateCommand(parsedComand.Command())
	resp_builder := domain.NewResponseBuilder()

	if command == CmdInvalid {
		resp_builder.Error("command: " + parsedComand.Command() + " " + error_msg.BAD_COMMAND)
		return nil, resp_builder.Build()
	}
	return nil, resp_builder.Build()
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
