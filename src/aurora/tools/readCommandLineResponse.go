package tools

import (
	"bufio"
	"fmt"
	"os"
)

// ReadCommandLineResponse read the response from the question
func ReadCommandLineResponse(question string) string {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Print(question + " ")
	text, err := reader.ReadString('\n')

	EasyPanic(err)

	return string(text)
}
