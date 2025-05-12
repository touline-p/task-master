package linereaders

import (
	"bufio"
	"fmt"
	"os"

	"github.com/touline-p/task-master/cli/domain"
)

type CliReader struct{}

func (clrdr *CliReader) Run() (string, domain.IResponse) {
	fmt.Print("Enter a line: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return line, nil
}

type SocketReader struct{}

func (clrdr *SocketReader) Run() (string, domain.IResponse) {
	fmt.Print("This is falsly a socket reader: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return line, nil
}
