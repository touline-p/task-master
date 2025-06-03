package linereaders

import (
	"bufio"
	"fmt"
	"os"

	"github.com/touline-p/task-master/cli/domain/interfaces"
	"github.com/touline-p/task-master/cli/domain"
)

type CliReader struct{}

func (clrdr *CliReader) Run() (string, interfaces.IResponse) {
	fmt.Print("Enter a line: ")
	reader := bufio.NewReader(os.Stdin)
	line, error := reader.ReadString('\n')

	if error != nil {
		resp_builder := domain.NewResponseBuilder()
		resp_builder.Error(error.Error())
		return line, resp_builder.Build()
	}
	return line, nil
}

type SocketReader struct{}

func (clrdr *SocketReader) Run() (string, interfaces.IResponse) {
	fmt.Print("This is falsly a socket reader: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return line, nil
}
