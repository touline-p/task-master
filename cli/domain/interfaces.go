package domain

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
	Code() CommandCode
	JobIds() []Job
}

type ISanitizer interface {
	Run(IParsedCommand) (ISanitizedCommand, IResponse)
}

type IStatusGetter interface {
	Run(Job) Status
}

type ILauncher interface {
	Run(ISanitizedCommand) IResponse
}

type IApplication interface {
	Run() (any, IResponse)
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
}
