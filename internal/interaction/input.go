package interaction

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hreluz/images-scrapper/pkg/html_processer/pagination"
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

func GetTagData(tagOptionText string, selectorTypeText string) *tag.Tag {

	ShowTagOptions(tagOptionText)

	tagName := GetTagChoice()

	ShowSelectorOptions(fmt.Sprintf("for the %s", tagName))

	selectorType := GetSelectorTypeChoice()

	if selectorType == selector.NONE {
		return tag.New(selector.Empty(), tagName)
	}

	return tag.New(
		selector.New(
			selectorType,
			GetUserInputWithErrorHandling(selectorTypeText),
		),
		tagName,
	)
}

func GetTagConfig(s string) *tag.TagConfig {
	input := GetUserInputWithErrorHandling(s)
	levels, _ := strconv.Atoi(input)
	tc := tag.NewConfig(levels, []tag.Tag{})

	for i := 0; i < levels; i++ {
		fmt.Println()

		t := GetTagData(
			fmt.Sprintf("for your level %v", (i+1)),
			"select your selector name",
		)
		fmt.Print("\n")
		tc.AddTag(t)
	}

	return tc
}

func GetPagination() *pagination.Pagination {
	resp := GetUserInputWithErrorHandling("Does this URL have pagination (Y/N)?")

	if resp == "Y" {
		v := GetUserInputWithErrorHandling("How many URLs would you like to check?")

		number, err := strconv.Atoi(v)

		if err != nil {
			log.Fatalf("Invalid number for pagination, error: %s", err)
		}

		paginationConfig := GetTagConfig("Insert how many levels the pagination will have, and where to get the next link: ")

		return pagination.New(
			paginationConfig,
			number,
		)
	}

	return pagination.New(nil, 1)
}
