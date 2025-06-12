package application

import (
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

	var startErrors []error
	for _, j := range jobs {
		if j.Config.AutoStart {
			startCmd := &cqrs.StartJobCommand{JobId: j.Id}
			if err := controller.CommandHandler().HandleStartJob(startCmd); err != nil {
				startErrors = append(startErrors, err)
			}
		}
	}

	if len(startErrors) > 0 {
		return services.ConcatenateErrors(startErrors)
	}

	return nil
}

func createDummyJob() models.Job {
	newEnv := make(map[string]string)
	newJobConfigValues := make([]models.JobConfigValue, 0, 1)
	newJobConfig := models.JobConfig{
		Name:          "Tail",
		Command:       "tail -f /dev/null",
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
	return *models.NewJob("Tail", newJobConfig)
}
