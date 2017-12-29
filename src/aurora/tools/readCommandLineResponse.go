package tools

import (
	"bufio"
	"cli/tools"
	"fmt"
	"os"
)

// ReadCommandLineResponse read the response from the question
func ReadCommandLineResponse(question string) string {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Print(question + " ")
	text, err := reader.ReadString('\n')

	tools.EasyPanic(err)

	return string(text)
}
