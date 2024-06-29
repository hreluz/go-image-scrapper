package interaction

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
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

func GetTagChoice() tag.TagName {

	optionsLength := len(tag.TAGS_OPTIONS)
	option := 1

	for {
		choice := GetUserInputWithErrorHandling("Your choice: ")
		option, _ = strconv.Atoi(choice)

		if option < 1 || option > optionsLength {
			continue
		}

		break
	}

	return tag.TAGS_OPTIONS[option-1]
}

func GetSelectorTypeChoice() selector.SelectorType {

	optionsLength := len(selector.SELECTOR_TYPE_OPTIONS)
	option := 1

	for {
		choice := GetUserInputWithErrorHandling("Your choice: ")
		option, _ = strconv.Atoi(choice)

		if option < 1 || option > optionsLength {
			continue
		}

		break
	}

	return selector.SELECTOR_TYPE_OPTIONS[option-1]
}
