package main

import (
	cli_services "github.com/touline-p/task-master/cli/domain/services"
	"github.com/touline-p/task-master/supervisor/application"
)

func main() {
	application.StartUpSupervisor()
	cli_services.InterpreteOneUserCommand()
}
