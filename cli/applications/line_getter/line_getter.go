package linegetter

import (
	"github.com/touline-p/task-master/cli/domain"
)

type SimpleLineGetter struct {}

func (lg *SimpleLineGetter) Run(readers []domain.IReader) (string, error) {
	return readers[0].Run()
}
