package interaction

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserInput(prompt string) (string, error) {
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

func GetUserInputWithErrorHandling(prompt string) string {
	input, err := getUserInput(prompt)
	if err != nil {
		log.Fatalf("There was an error when getting the input, error: %s", err)
	}
	return input
}
