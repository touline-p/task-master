package application

import (
	"context"
	"os"

	config_parser "github.com/touline-p/task-master/config_parser/domain"

	"github.com/touline-p/task-master/supervisor"
	"github.com/touline-p/task-master/supervisor/application/services"
	"github.com/touline-p/task-master/supervisor/domain/models"
)

func StartUpSupervisor() error {
	ctx := context.Background()
	controller := supervisor.GetSupervisorController()


	configs, err := config_parser.LoadConfig()
	if err != nil {
		return err
	}

	jobService := controller.JobService()

	jobs := make([]models.Job, len(*configs))
	for index, config := range(*configs) {
		jobs[index] = *models.NewJob(models.JobId(config.Name), config)
	}

	err = controller.Scheduler().RegisterJobs(jobs)

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
			if err := jobService.StartJob(ctx, j.Id); err != nil {
				startErrors = append(startErrors, err)
			}
		}
	}

	var stopErrors []error
	for _, j := range jobs {
		if j.Config.AutoStart {
			if err := jobService.StopJob(ctx, j.Id); err != nil {
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
