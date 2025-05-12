package services

type SupervisorService struct{}

func (ss *SupervisorService) Start() error {
	sc := SchedulerService{}
	error := sc.Initialize()
	if error != nil {
		return error
	}
	return nil
}
