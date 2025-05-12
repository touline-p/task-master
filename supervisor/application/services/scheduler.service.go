package services

import (
	"github.com/touline-p/task-master/supervisor"
	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
)

type SchedulerService struct {
	repository repositories.IJobRepository
}

func (ss *SchedulerService) Initialize() error {
	ss.repository = supervisor.GetControlerSupervisor().Repository()
	return nil
}

func (ss *SchedulerService) RegisterJobs(jobs []models.Job) error {
	errors := make([]error, 0, 1)
	for _, j := range jobs {
		err := ss.repository.Save(&j)
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

func (ss *SchedulerService) FindJob(id models.JobId) (models.Job, error) {
	return ss.repository.FindById(id)
}

func (ss *SchedulerService) FindAllJobs() ([]models.Job, error) {
	return ss.repository.FindAll()
}

func (ss *SchedulerService) LaunchJob(id models.JobId) error {
	job, retError := ss.FindJob(id)
	if retError != nil {
		return retError
	}

	retError = job.Start()
	if retError != nil {
		return retError
	}
	return nil
}
