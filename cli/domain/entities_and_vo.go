package domain

type CommandCode int
type JobId string
type Status int

// job
type Job struct {
	jobId  JobId
	status Status
}

func (j *Job) JobId() JobId   { return j.jobId }
func (j *Job) Status() Status { return j.status }

type Response struct {
	errors   []string
	warnings []string
	info     []string
}

type ResponseBuilder struct {
	errors   []string
	warnings []string
	info     []string
}

func (b *ResponseBuilder) Warning(warn string) { b.warnings = append(b.warnings, warn) }
func (b *ResponseBuilder) Error(error string)  { b.errors = append(b.errors, error) }
func (b *ResponseBuilder) Info(warn string)    { b.warnings = append(b.warnings, warn) }
func (b *ResponseBuilder) HandleCmd(errors error) {
	if errors != nil {
		b.Error(errors.Error())
	}
}
func (b *ResponseBuilder) HandleQuery(responses []string, errors []error) {
	for _, error := range errors {
		if error != nil {
			b.Error(error.Error())
		}
	}
	for _, response := range responses {
		b.Info(response)
	}
}

func (b *ResponseBuilder) Build() Response {
	return Response{errors: b.errors, warnings: b.warnings, info: b.info}
}

func NewResponseBuilder() ResponseBuilder {
	return ResponseBuilder{
		errors:   []string{},
		warnings: []string{},
		info:     []string{},
	}
}

func (r Response) Format() string {
	format := "Errors : "
	for _, e := range r.errors {
		format += "Err : " + e
	}
	for _, w := range r.warnings {
		format += "Err : " + w
	}
	return format
}
