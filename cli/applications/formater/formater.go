package formater

import (
	"strings"

	"github.com/touline-p/task-master/cli/domain/interfaces"
)

type SimpleFormater struct{}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func (f *SimpleFormater) Run(r interfaces.IResponse) string {
	return_string := []string{} 

	for _, info := range(r.Infos()) { return_string = append(return_string, info) }
	for _, info := range(r.Errors()) { return_string = append(return_string, info) }
	for _, info := range(r.Warnings()) { return_string = append(return_string, info) }
	return_string = append(return_string, "")
	return strings.Join(return_string, "\n")
}
