package linegetter

import (
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

type SimpleLineGetter struct{
	Readers []interfaces.IReader
}

func (lg *SimpleLineGetter) Run() (string, interfaces.IResponse) {
	return lg.Readers[0].Run()
}
