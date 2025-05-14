package cli_services

import cli_dependency "github.com/touline-p/task-master/cli"

func InterpreteOneUserCommand() {
	controler := cli_dependency.GetControlerCli()

	linegetter := controler.LineGetter()
	parser := controler.Parser()
	sanitizer := controler.Sanitizer()
	launcher := controler.Launcher()
	formater := controler.Formater()
	sender := controler.Sender()

	line, response := linegetter.Run(controler.Readers())
	parsedCommand, response := parser.Run(&line)
	sanitizedCommand, response := sanitizer.Run(parsedCommand)
	response = launcher.Run(sanitizedCommand)
	formatedString := formater.Run(response)
	sender.Run(formatedString)
}
