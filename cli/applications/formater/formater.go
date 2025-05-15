package formater

import "github.com/touline-p/task-master/cli/domain/interfaces"

type SimpleFormater struct{}

func (f *SimpleFormater) Run(r interfaces.IResponse) string {
	return ""
}
