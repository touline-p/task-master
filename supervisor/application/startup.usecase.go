package application

import (
	// TODO : Debug, a enlever
	"os"

	"github.com/touline-p/task-master/supervisor"
	"github.com/touline-p/task-master/supervisor/application/services"
	"github.com/touline-p/task-master/supervisor/domain/cqrs"
	"github.com/touline-p/task-master/supervisor/domain/models"
)

func StartUpSupervisor() error {
	controller := supervisor.GetSupervisorController()

	jobs := make([]models.Job, 0, 1)
	newJob := createDummyJob()
	jobs = append(jobs, newJob)

	err := controller.Scheduler().RegisterJobs(jobs)
	if err != nil {
		return err
	}

	jobs, err = controller.Repository().FindAll()
	if err != nil {
		return err
	}

	var errors []error
	for _, j := range jobs {
		startCmd := &cqrs.StartJobCommand{JobId: j.Id}
		if err := controller.CommandHandler().Handle(startCmd); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return services.ConcatenateErrors(errors)
	}

	return nil
}

func createDummyJob() models.Job {
	newExitCodes := make([]int, 0, 1)
	newEnv := make(map[string]string)
	newJobConfigValues := make([]models.JobConfigValue, 0, 1)
	newJobConfig := models.JobConfig{
		Name:          "Test Config",
		Command:       "Command",
		NumProcs:      1,
		Umask:         os.FileMode(os.O_RDONLY),
		WorkingDir:    "Working Dir",
		AutoStart:     true,
		AutoRestart:   models.RestartNever,
		ExitCodes:     newExitCodes,
		StartRetries:  0,
		StartDuration: 1,
		StopSignal:    "SIGKILL",
		StopDuration:  1,
		Stdout:        "",
		Stderr:        "",
		Environment:   newEnv,
		ConfigValues:  newJobConfigValues,
	}
	return *models.NewJob("Test job", newJobConfig)
}
