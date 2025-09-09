package application

import (
	"github.com/touline-p/task-master/supervisor"
)

func ShutdownSupervisor() error {
	controller := supervisor.GetSupervisorController()
	if err := controller.Scheduler().Shutdown(); err != nil {
		return err
	}
	return nil
}
