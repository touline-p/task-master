package application

import (
	"context"
	"os"

	"github.com/touline-p/task-master/supervisor"
	"github.com/touline-p/task-master/supervisor/application/services"
	"github.com/touline-p/task-master/supervisor/domain/models"
)

func StartUpSupervisor() error {
	ctx := context.Background()
	controller := supervisor.GetSupervisorController()

	jobService := controller.JobService()

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
		if j.Config.AutoStart {
			if err := jobService.StartJob(ctx, &j); err != nil {
				errors = append(errors, err)
			}
		}
	}

	if len(errors) > 0 {
		return services.ConcatenateErrors(errors)
	}

	return nil
}

func createDummyJob() models.Job {
	newEnv := make(map[string]string)
	newJobConfigValues := make([]models.JobConfigValue, 0, 1)
	newJobConfig := models.JobConfig{
		Name:          "Test Config",
		Command:       "echo 'Hello from Task Master'",
		NumProcs:      1,
		Umask:         os.FileMode(os.O_RDONLY),
		WorkingDir:    "",
		AutoStart:     true,
		AutoRestart:   models.RestartNever,
		ExitCodes:     []int{0},
		StartRetries:  2,
		StartDuration: 1,
		StopSignal:    "SIGKILL",
		StopDuration:  1,
		Stdout:        "/tmp/taskmaster.log",
		Stderr:        "/tmp/taskmaster-error.log",
		Environment:   newEnv,
		ConfigValues:  newJobConfigValues,
	}
	return *models.NewJob("Test job", newJobConfig)
}
