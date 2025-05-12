package launcher

import "github.com/touline-p/task-master/cli/domain"

type SimpleLauncher struct{}

type Response struct {
	content string
}

func (self *SimpleLauncher) Run(domain.ISanitizedCommand) domain.IResponse {
	builder := domain.NewResponseBuilder()
	return builder.Build()
}
