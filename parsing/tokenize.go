package parsing

import (
	"regexp"
	"strings"
)

func removeLineBreaks(content string) string {
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(content, " ")
}

func Tokenize(content string) []string {
	return strings.Split(content, " ")
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

func ParseFile(bytes []byte) (content string) {
	content = string(bytes)
	content = removeLineBreaks(content)
	content = removeMultipleSpaces(content)
	return
}
