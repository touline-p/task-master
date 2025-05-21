package ports

import (
	"context"

	"github.com/touline-p/task-master/supervisor/domain/models"
)

type ProcessManager interface {
	SpawnProcess(ctx context.Context, job *models.Job) (int, error)
}
