package application

import (
	// TODO : Debug, a enlever
	"os"

	"github.com/touline-p/task-master/supervisor/application/services"
	"github.com/touline-p/task-master/supervisor/domain/models"
)

func StartUpSupervisor() error {
	sup := services.SupervisorService{}
	retErr := sup.Start()
	if retErr != nil {
		return retErr
	}

	sch := services.SchedulerService{}
	retErr = sch.Initialize()
	if retErr != nil {
		return retErr
	}

	jobs := make([]models.Job, 0, 1)
	newJob := createDummyJob()
	jobs = append(jobs, newJob)
	retErr = sch.RegisterJobs(jobs)
	if retErr != nil {
		return retErr
	}

	jobs, retErr = sch.FindAllJobs()
	if retErr != nil {
		return retErr
	}

	errors := make([]error, 0, 1)
	for _, j := range jobs {
		err := sch.LaunchJob(j.Id)
		if err != nil {
			errors = append(errors, err)
		}
	}
	// TODO : Concatenate errors
	if len(errors) != 0 {
		return errors[1]
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
