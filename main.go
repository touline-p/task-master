package main

import (
	cli_services "github.com/touline-p/task-master/cli/domain/services"
	"github.com/touline-p/task-master/supervisor/application"
)

func main() {
	if err := application.StartUpSupervisor(); err != nil {
		panic(err)
	}
	for {
		cli_services.InterpreteOneUserCommand()
	}
}
