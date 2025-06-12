package infrastructure

import (
	"errors"
	"maps"
	"slices"
	"sync"

	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
)

var (
	ErrJobNotFound = errors.New("job not found")
	ErrNilJob      = errors.New("cannot save nil job")
)

type JobRepository struct {
	jobsMap map[models.JobId]models.Job
}

var lock = &sync.Mutex{}

var singleJobRepository *JobRepository

func GetJobRepository() repositories.IJobRepository {
	if singleJobRepository == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleJobRepository == nil {
			singleJobRepository = &JobRepository{jobsMap: map[models.JobId]models.Job{}}
		}
	}
	return singleJobRepository
}

func (jr JobRepository) Save(job *models.Job) error {
	if job == nil {
		return ErrNilJob
	}
	jr.jobsMap[job.Id] = *job
	return nil
}

func (jr JobRepository) FindById(id models.JobId) (models.Job, error) {
	job, exists := jr.jobsMap[id]
	if !exists {
		return models.Job{}, ErrJobNotFound
	}
	return job, nil
}

func (jr JobRepository) FindAll() ([]models.Job, error) {
	jobsList := maps.Values(jr.jobsMap)
	return slices.Collect(jobsList), nil
}

func (jr JobRepository) FindByStatus(status models.JobStatus) ([]models.Job, error) {
	var result []models.Job
	for _, job := range jr.jobsMap {
		if job.GetState().Status == status {
			result = append(result, job)
		}
	}
	return result, nil
}

func (jr JobRepository) Delete(id models.JobId) error {
	if !jr.Exists(id) {
		return ErrJobNotFound
	}
	delete(jr.jobsMap, id)
	return nil
}

func (jr JobRepository) Exists(id models.JobId) bool {
	_, exists := jr.jobsMap[id]
	return exists
}
