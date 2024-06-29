package interaction

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	htmlprocesser "github.com/hreluz/images-scrapper/pkg/html_processer"
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

func GetTagChoice() htmlprocesser.TagName {

	optionsLength := len(htmlprocesser.TAGS_OPTIONS)
	option := 1

	for {
		choice := GetUserInputWithErrorHandling("Your choice: ")
		option, _ = strconv.Atoi(choice)

		if option < 1 || option > optionsLength {
			continue
		}

		break
	}

	return htmlprocesser.TAGS_OPTIONS[option-1]
}

func GetSelectorTypeChoice() htmlprocesser.SelectorType {

	optionsLength := len(htmlprocesser.SELECTOR_TYPE_OPTIONS)
	option := 1

	for {
		choice := GetUserInputWithErrorHandling("Your choice: ")
		option, _ = strconv.Atoi(choice)

		if option < 1 || option > optionsLength {
			continue
		}

		break
	}

	return htmlprocesser.SELECTOR_TYPE_OPTIONS[option-1]
}
