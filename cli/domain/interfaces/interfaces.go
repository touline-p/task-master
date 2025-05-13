package interfaces

type ICliControler interface {
	Run()
	LineGetter() ILineGetter
	Parser()     IParser
	Sanitizer()  ISanitizer
	Launcher()   ILauncher
}

type IReader interface {
	Run() (string, IResponse)
}


type ILineGetter interface {
	Run() (string, IResponse)
}

type IParsedCommand interface {
	Command() string
	JobNames() []string
}

type IParser interface {
	Run(*string) (IParsedCommand, IResponse)
}

type ICommandCode interface {}

type IJob interface {}

type IStatus interface {}

type ISanitizedCommand interface {
	Code() ICommandCode
	JobIds() []IJob
}

type ISanitizer interface {
	Run(IParsedCommand) (ISanitizedCommand, IResponse)
}

type IStatusGetter interface {
	Run(IJob) IStatus
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
