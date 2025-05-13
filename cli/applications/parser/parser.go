package parser

import (
	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/cli/infrastructure/parsing"
	"github.com/touline-p/task-master/core/error_msg"
)

type SimpleParser struct{}

func (self *SimpleParser) Run(line *string) (interfaces.IParsedCommand, interfaces.IResponse) {
	words := parsing.SplitSpaces(line)
	resp_builder := domain.NewResponseBuilder()
	if len(words) == 0 {
		resp_builder.Error(error_msg.NO_INPUT)
		return nil, resp_builder.Build()
	}
	return &SimpleParsedCommand{
		command:  words[0],
		jobNames: words[1:],
	}, nil
}

type SimpleParsedCommand struct {
	command  string
	jobNames []string
}

func (c *SimpleParsedCommand) Command() string    { return c.command }
func (c *SimpleParsedCommand) JobNames() []string { return c.jobNames }
