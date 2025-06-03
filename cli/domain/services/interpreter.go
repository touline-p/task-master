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
	print(line)
	parsedCommand, response := parser.Run(&line)
	print(parsedCommand)
	sanitizedCommand, response := sanitizer.Run(parsedCommand)
	print(sanitizedCommand)
	response = launcher.Run(sanitizedCommand)
	print(response)
	formatedString := formater.Run(response)
	print(formatedString)
	sender.Run(formatedString)
}
