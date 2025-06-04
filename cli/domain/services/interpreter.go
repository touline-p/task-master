package cli_services

import (
	cli_dependency "github.com/touline-p/task-master/cli"
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

func InterpreteOneUserCommand() {
	controler := cli_dependency.GetControlerCli()


	formater := controler.Formater()
	sender := controler.Sender()

	response := readAndExecuteLine()
	formatedString := formater.Run(response)
	sender.Run(formatedString)
}

func readAndExecuteLine() interfaces.IResponse {
	controler := cli_dependency.GetControlerCli()

	linegetter := controler.LineGetter()
	parser := controler.Parser()
	sanitizer := controler.Sanitizer()
	launcher := controler.Launcher()

	line, response := linegetter.Run(controler.Readers())
	if response != nil { return response }

	parsedCommand, response := parser.Run(&line)
	if response != nil { return response }

	sanitizedCommand, response := sanitizer.Run(parsedCommand)
	if response != nil { return response }

	response = launcher.Run(sanitizedCommand)
	return response
}
