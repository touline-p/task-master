package linereaders

import (
	"bufio"
	"fmt"
	"os"
)

type CliReader struct {}

func (clrdr *CliReader)Run() (string, error) {
	fmt.Print("Enter a line: ")
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

type SocketReader struct {}

func (clrdr *SocketReader)Run() (string, error) {
	fmt.Print("This is falsly a socket reader: ")
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
