package utils

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func removeLineBreaks(content string) string {
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(content, " ")
}

func addSpacesToBrackets(content string) string {
	addSpacesToBrackets := regexp.MustCompile(`(\(|\))`)
	return strings.TrimSpace(
		removeMultipleSpaces(
			addSpacesToBrackets.ReplaceAllString(content, " $1 "),
		),
	)
}

func removeMultipleSpaces(content string) string {
	removeMultipleSpaces := regexp.MustCompile(`\s+`)
	return removeMultipleSpaces.ReplaceAllString(content, " ")
}

func ReadFile(fileName string) (content string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content = string(bytes)
	content = removeLineBreaks(content)
	content = removeMultipleSpaces(content)
	return
}

func SplitTokens(content string) []string {
	//"( if ( > 2 3 ) )"
	return strings.Split(content, " ")
}
