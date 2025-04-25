package cli

import (
	launcher "github.com/touline-p/task-master/cli/applications/launcher"
	linegetter "github.com/touline-p/task-master/cli/applications/line_getter"
	parser "github.com/touline-p/task-master/cli/applications/parser"
	sanitizer "github.com/touline-p/task-master/cli/applications/sanitizer"
	"github.com/touline-p/task-master/cli/domain"
	linereaders "github.com/touline-p/task-master/cli/infrastructure/line_readers"
)

type Controler struct {
	readers    []domain.IReader
	lineGetter domain.ILineGetter
	parser     domain.IParser
	sanitizer  domain.ISanitizer
	launcher   domain.ILauncher
}

func (c *Controler) Readers() []domain.IReader      { return c.readers }
func (c *Controler) LineGetter() domain.ILineGetter { return c.lineGetter }
func (c *Controler) Parser() domain.IParser         { return c.parser }
func (c *Controler) Sanitizer() domain.ISanitizer   { return c.sanitizer }
func (c *Controler) Launcher() domain.ILauncher     { return c.launcher }

func GetControlerCli() domain.IControler {
	return &Controler{
		readers: []domain.IReader{
			&linereaders.CliReader{},
			&linereaders.SocketReader{},
		},
		lineGetter: &linegetter.SimpleLineGetter{},
		parser:     &parser.SimpleParser{},
		sanitizer:  &sanitizer.SimpleSanitizer{},
		launcher:   &launcher.SimpleLauncher{},
	}
}
