package main

import (
	"github.com/touline-p/task-master/supervisor/application"
	"github.com/touline-p/task-master/dependency_injection"
)

func main() {
	var controler dependency_injection.IControler
	application.StartUpSupervisor()

	controler = dependency_injection.GetSimpleControler()
	cli_controler := controler.CliEntryPoint()
	cli_controler.Run()
}
