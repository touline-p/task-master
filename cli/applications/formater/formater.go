package formater

import (
	"fmt"
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

	for _, info := range r.Infos() {
		return_string = append(return_string, info)
	}
	for _, info := range r.Errors() {
		return_string = append(return_string, info)
	}
	for _, info := range r.Warnings() {
		return_string = append(return_string, info)
	}
	return_string = append(return_string, "")
	return strings.Join(return_string, "\n")
}

type FancyFormater struct{}

func (f *FancyFormater) Run(r interfaces.IResponse) string {
	formated_infos := make([]string, len(r.Infos()))
	formated_warns := make([]string, len(r.Warnings()))
	formated_errors := make([]string, len(r.Errors()))
	for i, s := range r.Infos() {
		formated_infos[i] = fmt.Sprintf("%s%sINFO : %s%s", Reset, Gray, s, Reset)
	}
	for i, s := range r.Warnings() {
		formated_warns[i] = fmt.Sprintf("%s%sWARN : %s%s", Reset, Yellow, s, Reset)
	}
	for i, s := range r.Errors() {
		formated_errors[i] = fmt.Sprintf("%s%sERR  : %s%s", Reset, Red, s, Reset)
	}

	return fmt.Sprintf(
		"%s%s%s",
		strings.Join(formated_infos, ""),
		strings.Join(formated_warns, ""),
		strings.Join(formated_errors, ""),
	)
}
