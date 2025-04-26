package domain

type CommandCode int
type JobId string
type Status int


type IParsedCommand interface {
	Command() string
	JobNames() []string
}

type ISanitizedCommand interface {
	Code() CommandCode
	JobIds() []Job
}

// Response
type Response interface {
	FormatedString() string
}

// job
type Job struct {
	Id JobId
	Status Status
}
