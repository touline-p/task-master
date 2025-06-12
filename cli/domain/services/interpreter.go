package cli_services

import (
	cli_dependency "github.com/touline-p/task-master/cli"
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

func InterpreteOneUserCommand() {
	controler := cli_dependency.GetControlerCli()

	formater := controler.Formater()
	ioManager := controler.IOManager()

	resp_builder := readAndExecuteLine()
	formatedString := formater.Run(resp_builder)
	ioManager.Write(formatedString)
}

func readAndExecuteLine() interfaces.IResponse {
	controler := cli_dependency.GetControlerCli()

	ioManager := controler.IOManager()
	parser := controler.Parser()
	sanitizer := controler.Sanitizer()
	launcher := controler.Launcher()

	line, resp_builder := ioManager.Read()
	if resp_builder.HasErrors() {
		return resp_builder.Build()
	}

	parsedCommand, resp_builder := parser.Run(&line, resp_builder)
	if resp_builder.HasErrors() {
		return resp_builder.Build()
	}

	sanitizedCommand, resp_builder := sanitizer.Run(parsedCommand, resp_builder)
	if resp_builder.HasErrors() {
		return resp_builder.Build()
	}

	resp_builder = launcher.Run(sanitizedCommand, resp_builder)
	return resp_builder.Build()
}
