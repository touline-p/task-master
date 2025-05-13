package use_case

import (
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

type  SimpleCliEntryPoint struct {
	lineGetter interfaces.ILineGetter
	parser     interfaces.IParser
	sanitizer  interfaces.ISanitizer
	launcher   interfaces.ILauncher
}

func NewSimpleCliEntryPoint(
	lineGetter interfaces.ILineGetter,
	parser     interfaces.IParser,
	sanitizer  interfaces.ISanitizer,
	launcher   interfaces.ILauncher,
) *SimpleCliEntryPoint {
	return  &SimpleCliEntryPoint{
		lineGetter: lineGetter,
		parser: parser,
		sanitizer: sanitizer,
		launcher: launcher,
	}
}

func (c *SimpleCliEntryPoint) LineGetter() interfaces.ILineGetter { return c.lineGetter }
func (c *SimpleCliEntryPoint) Parser() interfaces.IParser         { return c.parser }
func (c *SimpleCliEntryPoint) Sanitizer() interfaces.ISanitizer   { return c.sanitizer }
func (c *SimpleCliEntryPoint) Launcher() interfaces.ILauncher     { return c.launcher }

func (ep *SimpleCliEntryPoint)Run() {
	linegetter := ep.lineGetter
	parser := ep.parser
	sanitizer := ep.sanitizer
	launcher := ep.launcher
	for {
		line, response := linegetter.Run()
		parsedCommand, response := parser.Run(&line)
		sanitizedCommand, response := sanitizer.Run(parsedCommand)
		response = launcher.Run(sanitizedCommand)
		println(response.Format())
	}
}
