package services

import (
	"context"
	"os"

	"github.com/touline-p/task-master/supervisor/domain/models"
)

// Low-level process operations
type IProcessManager interface {
	Launch(ctx context.Context, job *models.Job) (int, error)
	Signal(job *models.Job, signal os.Signal) error
	Terminate(job *models.Job) error
	Kill(job *models.Job) error
	WaitForExit(jobId models.JobId) (int, error, bool)
	// IsRunning(job *models.Job) bool
	// GetProcessInfo(job *models.Job) (*os.Process, error)
}
