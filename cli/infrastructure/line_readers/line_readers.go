package linereaders

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chzyer/readline"

	"github.com/touline-p/task-master/cli/applications/sanitizer"
	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

type CliManager struct {
}

func (clrdr *CliManager) Read() (string, interfaces.IResponseBuilder) {
	respBuilder := domain.NewResponseBuilder()
	fmt.Print("Enter a line: ")
	reader := bufio.NewReader(os.Stdin)
	line, error := reader.ReadString('\n')
	if error != nil {
		respBuilder.Error(fmt.Sprintln("Cli reader : ", error.Error()))
	}
	return line, respBuilder
}

func (clrdr *CliManager) Write(formatedResponse string) {
	println(formatedResponse)
}

type AutoCompleteManager struct {
}

func (clrdr *AutoCompleteManager) Read() (string, interfaces.IResponseBuilder) {
	respBuilder := domain.NewResponseBuilder()

	var items []readline.PrefixCompleterInterface
	for _, s := range sanitizer.CommandSymbols {
		items = append(items, readline.PcItem(s))
	}

	completer := readline.NewPrefixCompleter(
		readline.PcItem("command", items...),
	)

	rl, _ := readline.NewEx(
		&readline.Config{Prompt: "task-master - > ", AutoComplete: completer},
	)
	defer rl.Close()

	line, err := rl.Readline()
	if err == nil {
		os.Exit(1)
	}
	return line, respBuilder
}

func (clrdr *AutoCompleteManager) Write(formatedResponse string) {
	println(formatedResponse)
}
