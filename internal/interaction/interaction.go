package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetUserInput(prompt string) (string, error) {
	if prompt != "" {
		fmt.Printf("%s: ", prompt)
	}

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput, nil
}
