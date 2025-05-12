package supervisor

import (
	"github.com/touline-p/task-master/supervisor/domain/repositories"
	"github.com/touline-p/task-master/supervisor/infrastructure"
)

type Controler struct {
	repository repositories.IJobRepository
}

func (c *Controler) Repository() repositories.IJobRepository { return c.repository }

func GetControlerSupervisor() *Controler {
	return &Controler{
		repository: infrastructure.NewJobRepository(),
	}
}
