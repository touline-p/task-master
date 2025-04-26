package sanitizer

import (
	"errors"

	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleSanitizer struct {}

type SanitizedCommand struct {
	command domain.CommandCode
	jobs []domain.Job
}


func (c *SanitizedCommand) Command() domain.CommandCode { return c.command }
func (c *SanitizedCommand) JobNames() []domain.Job { return c.jobs }


func (self *SimpleSanitizer)Run(parsedComand domain.IParsedCommand) (domain.ISanitizedCommand, error) {
	command := ValidateCommand(parsedComand.Command())
	if command == CmdInvalid {
		return nil, errors.New(error_msg.BAD_COMMAND)
	}
	jobs := getJobStatus
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

var CommandSymbols = [CmdNumber]string {
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
