package use_case

import (
	"fmt"
	"github.com/touline-p/task-master/cli"
)

func Run ()  {
	fmt.Println("bonjour")
	controler := cli.GetControlerCli()
	line, _ := controler.LineGetter().Run(controler.Readers())
	parsedCommand, _ := controler.Parser().Run(line)
	sanitizedCommand, _ := controler.Sanitizer().Run(parsedCommand)
	response, _ := controler.Launcher().Run(sanitizedCommand)
	println(response.Format())
}
