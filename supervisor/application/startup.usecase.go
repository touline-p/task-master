package application

import (
	"context"
	"os"
	"time"

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

	var startErrors []error
	for _, j := range jobs {
		if j.Config.AutoStart {
			if err := jobService.StartJob(ctx, &j); err != nil {
				startErrors = append(startErrors, err)
			}
		}
	}

	time.Sleep(10 * time.Second)

	var stopErrors []error
	for _, j := range jobs {
		if j.Config.AutoStart {
			if err := jobService.StopJob(ctx, &j); err != nil {
				stopErrors = append(stopErrors, err)
			}
		}
	}

	if len(startErrors) > 0 {
		return services.ConcatenateErrors(startErrors)
	}

	if len(stopErrors) > 0 {
		return services.ConcatenateErrors(stopErrors)
	}

	return nil
}

func createDummyJob() models.Job {
	newEnv := make(map[string]string)
	newJobConfigValues := make([]models.JobConfigValue, 0, 1)
	newJobConfig := models.JobConfig{
		Name:          "Yes",
		Command:       "yes",
		NumProcs:      1,
		Umask:         os.FileMode(os.O_RDONLY),
		WorkingDir:    "",
		AutoStart:     true,
		AutoRestart:   models.RestartNever,
		ExitCodes:     []int{0},
		StartRetries:  2,
		StartDuration: 1,
		StopSignal:    "SIGTERM",
		StopDuration:  1,
		Stdout:        "/tmp/taskmaster.log",
		Stderr:        "/tmp/taskmaster-error.log",
		Environment:   newEnv,
		ConfigValues:  newJobConfigValues,
	}
	return *models.NewJob("Yes", newJobConfig)
}
