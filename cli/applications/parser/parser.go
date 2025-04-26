package parser

import (
	"errors"

	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/infrastructure/parsing"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleParser struct {}

func (self *SimpleParser)Run(line *string) (domain.IParsedCommand, error) {
	words := parsing.SplitSpaces(line)
	if len(words) == 0 {
		return nil, errors.New(error_msg.NO_INPUT)
	}
	return  &SimpleParsedCommand{
		command: words[0],
		jobNames: words[1:],
	}, nil
}

type SimpleParsedCommand struct {
	command string
	jobNames []string
}

func (c *SimpleParsedCommand) Command() string { return c.command }
func (c *SimpleParsedCommand) JobNames() []string { return c.jobNames }
