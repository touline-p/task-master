package interfaces

type IIOManager interface {
	Read() (string, IResponseBuilder)
	Write(string)
}

type IParsedCommand interface {
	Command() string
	JobNames() []string
}

type IResponseBuilder interface {
	Build() IResponse
	Info(string)
	Warning(string)
	Error(string)
	HasErrors() bool
}

type IParser interface {
	Run(*string, IResponseBuilder) (IParsedCommand, IResponseBuilder)
}

type ISanitizedCommand interface {
	Code() ICommandCode
	JobIds() []IJob
}
type ICommandCode interface{}
type IJob interface {
	Id() string
}
type IStatus interface{}

type ISanitizer interface {
	Run(IParsedCommand, IResponseBuilder) (ISanitizedCommand, IResponseBuilder)
}

type IStatusGetter interface {
	Run(IJob) IStatus
}

type ILauncher interface {
	Run(ISanitizedCommand, IResponseBuilder) IResponseBuilder
	SvTranslator() ISupervisorTranslator
	SvAdapter() ISupervisorAdapter
}

type IFormater interface {
	Run(IResponse) string
}

type ISender interface {
	Run(string)
}

type IResponse interface {
	Infos() []string
	Errors() []string
	Warnings() []string
}

type IControler interface {
	IOManager() IIOManager
	Parser() IParser
	Sanitizer() ISanitizer
	Launcher() ILauncher
	Formater() IFormater
	Sender() ISender
}
