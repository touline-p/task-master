package cqrs

import (
	"github.com/touline-p/task-master/supervisor/domain/models"
)

type Command interface {
	Type() string
}

type StartJobCommand struct {
	JobId models.JobId
}

func (c *StartJobCommand) Type() string {
	return "START_JOB"
}

type StopJobCommand struct {
	JobId models.JobId
}

func (c *StopJobCommand) Type() string {
	return "STOP_JOB"
}

type RestartJobCommand struct {
	JobId models.JobId
}

func (c *RestartJobCommand) Type() string {
	return "RESTART_JOB"
}

type ReloadConfigCommand struct {
	ConfigPath string
}

func (c *ReloadConfigCommand) Type() string {
	return "RELOAD_CONFIG"
}

type StopSupervisorCommand struct{}

func (c *StopSupervisorCommand) Type() string {
	return "STOP_SUPERVISOR"
}

type ProcessEventCommand struct {
	Event models.ProcessEvent
}

func (c *ProcessEventCommand) Type() string {
	return "PROCESS_EVENT"
}

type ICommandHandler interface {
	Handle(command Command) error
}
