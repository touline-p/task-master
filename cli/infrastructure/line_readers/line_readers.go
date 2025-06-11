package linereaders

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/touline-p/task-master/cli/domain"
	"github.com/touline-p/task-master/cli/domain/interfaces"
)

type CliManager struct{
	instream io.Reader
	outstream io.Writer
	closer io.Closer
}


func (clrdr *CliManager) Read() (string, interfaces.IResponse) {
	fmt.Print("Enter a line: ")
	reader := bufio.NewReader(os.Stdin)
	line, error := reader.ReadString('\n')
	if error != nil {
		resp_builder := domain.NewResponseBuilder()
		resp_builder.Error(fmt.Sprintln("Cli reader : ", error.Error()))
		return line, resp_builder.Build()
	}
	return line, nil
}

func (clrdr *CliManager) Write(formated_response string) () {
	println(formated_response);
}
