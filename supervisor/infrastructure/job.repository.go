package infrastructure

import (
	"maps"
	"slices"

	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/repositories"
)

type JobRepository struct {
	jobsMap map[models.JobId]models.Job
}

func NewJobRepository() repositories.IJobRepository {
	return &JobRepository{jobsMap: map[models.JobId]models.Job{}}
}

func (jr JobRepository) Save(job *models.Job) error {
	// if job == nil {
	// 	return CannotSaveError
	// }
	jr.jobsMap[job.Id] = *job
	return nil
}

func (jr JobRepository) FindById(id models.JobId) (models.Job, error) {
	job := jr.jobsMap[id]
	return job, nil
}

func (jr JobRepository) FindAll() ([]models.Job, error) {
	jobsList := maps.Values(jr.jobsMap)
	return slices.Collect(jobsList), nil
}
