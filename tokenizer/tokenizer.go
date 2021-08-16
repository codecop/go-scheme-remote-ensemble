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

func Scan(input string) ([]Token, error) {
	if len(input) == 0 {
		return nil, nil
	}

	terms := createTerms(input)
	tokens := createTokens(terms)

	if len(tokens) > 0 {
		return tokens, nil
	} else {
		return nil, fmt.Errorf("no valid token %s", input)
	}
}

func createTerms(input string) []string {
	addSpacesToBrackets := regexp.MustCompile(`(\(|\))`)
	spacedInput := addSpacesToBrackets.ReplaceAllString(input, " $1 ")
	return strings.Fields(spacedInput)
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
 * Wording: parents (), edged brackets <>, square brackets [], curly braces {}
 */
