package main

import (
	"github.com/touline-p/task-master/supervisor/application"
)

func main() {
	if err := application.StartUpSupervisor(); err != nil {
		println(err)
	}
	// for {
	// 	cli_services.InterpreteOneUserCommand()
	// }
	if err := application.ShutdownSupervisor(); err != nil {
		panic(err)
	}
}
