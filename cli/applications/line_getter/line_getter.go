package linegetter

import (
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

type SimpleLineGetter struct{}

func (lg *SimpleLineGetter) Run(readers []interfaces.IReader) (string, interfaces.IResponse) {
	return readers[0].Run()
}
