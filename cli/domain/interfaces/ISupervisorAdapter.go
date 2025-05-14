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
	StopJob(jobId string) error
	RestartJob(jobId string) error
	GetJobStatuses() (map[string]JobStatus, error)
}
