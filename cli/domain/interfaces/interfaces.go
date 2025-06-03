package interfaces

type IReader interface {
	Run() (string, IResponse)
}

type ILineGetter interface {
	Run([]IReader) (string, IResponse)
}

type IParsedCommand interface {
	Command() string
	JobNames() []string
}

type IParser interface {
	Run(*string) (IParsedCommand, IResponse)
}

type ISanitizedCommand interface {
	Code() ICommandCode
	JobIds() []IJob
}
type ICommandCode interface{}
type IJob interface {
	ToString() string
}
type IStatus interface{}

type ISanitizer interface {
	Run(IParsedCommand) (ISanitizedCommand, IResponse)
}

type IStatusGetter interface {
	Run(IJob) IStatus
}

type ILauncher interface {
	Run(ISanitizedCommand) IResponse
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
	Format() string
}

type IControler interface {
	Readers() []IReader
	LineGetter() ILineGetter
	Parser() IParser
	Sanitizer() ISanitizer
	Launcher() ILauncher
	Formater() IFormater
	Sender() ISender
}
