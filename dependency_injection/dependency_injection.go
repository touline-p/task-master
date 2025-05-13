package dependency_injection

import (
	"github.com/touline-p/task-master/cli/applications/launcher"
	"github.com/touline-p/task-master/cli/applications/line_getter"
	"github.com/touline-p/task-master/cli/applications/parser"
	"github.com/touline-p/task-master/cli/applications/entrypoint"
	"github.com/touline-p/task-master/cli/applications/sanitizer"
	cli_interfaces "github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/cli/infrastructure/line_readers"
)


type IControler interface {
	CliEntryPoint() cli_interfaces.ICliControler
}

type SimpleControler struct {
	cliEntryPoint cli_interfaces.ICliControler
}

func (sc * SimpleControler)CliEntryPoint() cli_interfaces.ICliControler{ return sc.cliEntryPoint }



func GetSimpleControler()  IControler {
	return &SimpleControler {
		cliEntryPoint: use_case.NewSimpleCliEntryPoint(
			&linegetter.SimpleLineGetter{
				Readers: []cli_interfaces.IReader{
					&linereaders.CliReader{},
					&linereaders.SocketReader{},
				},
			},
			&parser.SimpleParser{},
			&sanitizer.SimpleSanitizer{},
			&launcher.SimpleLauncher{},
		),
	}
}
