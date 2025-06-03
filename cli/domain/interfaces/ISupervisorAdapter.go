package interfaces

type JobStatus int

const (
	_ JobStatus = iota
	StatusStopped
	StatusStarting
	StatusStopping
	StatusBackoff
	StatusRunning
	StatusExited
	StatusFatal
)

type ISupervisorAdapter interface {
	StartJob(jobId string) error
	StartJobs(jobIds []string) error
	StopJob(jobId string) error
	StopJobs(jobIds []string) error
	RestartJob(jobId string) error
	RestartJobs(jobIds []string) error
	GetJobStatuses(jobIds []string) (map[string]JobStatus, error)
}

type ISupervisorTranslator interface {
	Translate([]string) []string
}

