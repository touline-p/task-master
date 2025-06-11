package services

import (
	"os"

	"github.com/touline-p/task-master/supervisor/domain/models"
)

// Low-level process operations
type IProcessManager interface {
	Launch(job *models.Job) (int, error)
	Signal(job *models.Job, signal os.Signal) error
	Terminate(job *models.Job) error
	Kill(job *models.Job) error
	IsRunning(job *models.Job) bool
	GetProcessInfo(job *models.Job) (*os.Process, error)
}
