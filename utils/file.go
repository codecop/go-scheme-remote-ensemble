package utils

import (
	"io/ioutil"
	"log"
	"regexp"
)

func removeLineBreaks(content string) string {
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(content, " ")
}

func removeSpaces(content string) string {
	re := regexp.MustCompile(`\s+`)
	openingBracketSpaces := regexp.MustCompile(` \( |\( | \(`)
	closingBracketSpaces := regexp.MustCompile(` \) |\) | \)`)
	return closingBracketSpaces.ReplaceAllString(
		openingBracketSpaces.ReplaceAllString(
			re.ReplaceAllString(content, " "),
			"("),
		")")
}

func ReadFile(fileName string) (content string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content = string(bytes)
	content = removeLineBreaks(content)
	content = removeSpaces(content)
	return
}
