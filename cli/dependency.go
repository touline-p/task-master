package cli

import (
	"github.com/touline-p/task-master/cli/applications/formater"
	launcher "github.com/touline-p/task-master/cli/applications/launcher"
	linegetter "github.com/touline-p/task-master/cli/applications/line_getter"
	parser "github.com/touline-p/task-master/cli/applications/parser"
	sanitizer "github.com/touline-p/task-master/cli/applications/sanitizer"
	"github.com/touline-p/task-master/cli/applications/sender"
	"github.com/touline-p/task-master/cli/domain/interfaces"
	linereaders "github.com/touline-p/task-master/cli/infrastructure/line_readers"
)

type Controler struct {
	readers    []interfaces.IReader
	lineGetter interfaces.ILineGetter
	parser     interfaces.IParser
	sanitizer  interfaces.ISanitizer
	launcher   interfaces.ILauncher
	formater   interfaces.IFormater
	sender     interfaces.ISender
}

func (c *Controler) Readers() []interfaces.IReader      { return c.readers }
func (c *Controler) LineGetter() interfaces.ILineGetter { return c.lineGetter }
func (c *Controler) Parser() interfaces.IParser         { return c.parser }
func (c *Controler) Sanitizer() interfaces.ISanitizer   { return c.sanitizer }
func (c *Controler) Launcher() interfaces.ILauncher     { return c.launcher }
func (c *Controler) Formater() interfaces.IFormater     { return c.formater }
func (c *Controler) Sender() interfaces.ISender         { return c.sender }

func GetControlerCli() interfaces.IControler {
	return &Controler{
		readers: []interfaces.IReader{
			&linereaders.CliReader{},
			&linereaders.SocketReader{},
		},
		lineGetter: &linegetter.SimpleLineGetter{},
		parser:     &parser.SimpleParser{},
		sanitizer:  &sanitizer.SimpleSanitizer{},
		launcher:   &launcher.SimpleLauncher{},
		formater:   &formater.SimpleFormater{},
		sender:     &sender.SimpleSender{},
	}
}
