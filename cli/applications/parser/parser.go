package parser

import (
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/cli/infrastructure/parsing"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleParser struct{}

func (self *SimpleParser) Run(line *string, bldr interfaces.IResponseBuilder) (interfaces.IParsedCommand, interfaces.IResponseBuilder) {
	words := parsing.SplitSpaces(line)
	if len(words) == 0 {
		bldr.Error(error_msg.NO_INPUT)
		return nil, bldr
	}
	return &SimpleParsedCommand{
		command:  words[0],
		jobNames: words[1:],
	}, bldr
}

type SimpleParsedCommand struct {
	command  string
	jobNames []string
}

func (c *SimpleParsedCommand) Command() string    { return c.command }
func (c *SimpleParsedCommand) JobNames() []string { return c.jobNames }
