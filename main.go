package main

import (
	"github.com/touline-p/task-master/supervisor/application"
	"github.com/touline-p/task-master/dependency_injection"
)

func main() {
	var controler dependency_injection.IControler
	application.StartUpSupervisor()

	controler = dependency_injection.GetSimpleControler()
	cli_controler := controler.CliEntryPoint()
	cli_controler.Run()
}

/*
Run() {
	line = reader()
	parser
	sanitizer
	launcher

}

domain:
	cli:
		dependency_injection.go
		domaine:
			typage Entite et value object
			interface des services
			const (
				CmdInvalid CommandCode = iota
				CmdExit
				CmdPid
				CmdHelp
				CmdUpdate
				CmdStart
				CmdStop
				CmdRestart
				CmdStatus
			)


		applications:
			use_case -> IParser
			parsing
			terminalReader
			networkReader

		infrastructure: (ReadLine) depend de ce qui est utilise techniquement
			read, recv

			repository:
				permet d'aller chercher les donnees
				l'interaction avec le stockage/retrieve data

		presentation:
			comment interagire avec le context

	supervisor:
*/
