package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

type Token interface {
}

type Scanner struct {
	IsToken  func(token string) bool
	NewToken func(token string) Token
}

var scanners = []Scanner{
	NewBoolScanner(),
	NewNumberScanner(),
	NewParenthesisScanner(),
	NewNameScanner(),
}

func Scan(cleanedString string) ([]Token, error) {
	if len(cleanedString) == 0 {
		return nil, nil
	}

	terms := createTermsFrom(cleanedString)
	tokens := createTokens(terms)

	if len(tokens) > 0 {
		return tokens, nil
	} else {
		return nil, fmt.Errorf("no valid token %s", cleanedString)
	}

}

func createTermsFrom(cleanedString string) []string {
	addSpacesToBrackets := regexp.MustCompile(`(\(|\))`)

	spacedString := addSpacesToBrackets.ReplaceAllString(cleanedString, " $1 ")
	terms := strings.Fields(spacedString)
	return terms
}

func createTokens(terms []string) []Token {
	tokens := []Token{}

	for _, term := range terms {
		for _, scanner := range scanners {
			if scanner.IsToken(term) {
				tokens = append(tokens, scanner.NewToken(term))
			}
		}
	}
	return tokens
}

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */

/**
* parents (), edged brackets <>, square brackets [], curly braces {}
 */
