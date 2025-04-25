package supervisor

type ISchedulerService interface {
	Jobs() []Job
	Job(id string) Job
	addJob(Job)
	startJob(Job)
	restartJob(Job)
	stopJob(Job)
}
