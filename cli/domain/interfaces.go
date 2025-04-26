package domain

type IReader interface {
	Run() (string, error)
}

type ILineGetter interface {
	Run([]IReader) (string, error)
}

type IParser interface {
	Run(string) (IParsedCommand, error)
}

type ISanitizer interface {
	Run(IParsedCommand) (ISanitizedCommand, error)
}

type IStatusGetter interface {
	Run(Job) Status
}

type ILauncher interface {
	Run(ISanitizedCommand) (IResponse, error)
}

type IApplication interface {
	Run() (any, error)
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
