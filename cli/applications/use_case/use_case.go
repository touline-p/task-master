package use_case

import (
	"github.com/touline-p/task-master/cli"
)

func Run() {
	controler := cli.GetControlerCli()
	linegetter := controler.LineGetter()
	parser := controler.Parser()
	sanitizer := controler.Sanitizer()
	launcher := controler.Launcher()
	for {
		line, response := linegetter.Run(controler.Readers())
		parsedCommand, response := parser.Run(&line)
		sanitizedCommand, response := sanitizer.Run(parsedCommand)
		response = launcher.Run(sanitizedCommand)
		println(response.Format())
	}
}
