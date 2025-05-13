package launcher

import (
	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

type SimpleLauncher struct{}

type Response struct {
	content string
}

func (self *SimpleLauncher) Run(interfaces.ISanitizedCommand) interfaces.IResponse {
	builder := domain.NewResponseBuilder()
	return builder.Build()
}
