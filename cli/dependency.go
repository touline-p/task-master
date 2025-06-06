package cli

import (
	"github.com/touline-p/task-master/cli/applications/formater"
	"github.com/touline-p/task-master/cli/applications/launcher"
	"github.com/touline-p/task-master/cli/applications/parser"
	"github.com/touline-p/task-master/cli/applications/sanitizer"
	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/cli/infrastructure"
	linereaders "github.com/touline-p/task-master/cli/infrastructure/line_readers"
	"github.com/touline-p/task-master/supervisor"
)

type Controler struct {
	iomanager interfaces.IIOManager
	parser    interfaces.IParser
	sanitizer interfaces.ISanitizer
	launcher  interfaces.ILauncher
	formater  interfaces.IFormater
	sender    interfaces.ISender
}

func (c *Controler) IOManager() interfaces.IIOManager { return c.iomanager }
func (c *Controler) Parser() interfaces.IParser       { return c.parser }
func (c *Controler) Sanitizer() interfaces.ISanitizer { return c.sanitizer }
func (c *Controler) Launcher() interfaces.ILauncher   { return c.launcher }
func (c *Controler) Formater() interfaces.IFormater   { return c.formater }
func (c *Controler) Sender() interfaces.ISender       { return c.sender }

func GetControlerCli() interfaces.IControler {
	controler := supervisor.GetSupervisorController()
	qryHdlr := controler.QueryHandler()
	cmdHdlr := controler.CommandHandler()
	return &Controler{
		iomanager: &linereaders.CliManager{},
		parser:    &parser.SimpleParser{},
		sanitizer: &sanitizer.SimpleSanitizer{
			SupervisorAdapter: infrastructure.NewSupervisorAdapter(
				cmdHdlr,
				qryHdlr,
			),
			SupervisorTranslator: &infrastructure.SupervisorTranslator{},
		},
		launcher: &launcher.SimpleLauncher{
			SupervisorAdapter: infrastructure.NewSupervisorAdapter(
				cmdHdlr,
				qryHdlr,
			),
			SupervisorTranslator: &infrastructure.SupervisorTranslator{},
		},
		formater: &formater.FancyFormater{},
	}
}
