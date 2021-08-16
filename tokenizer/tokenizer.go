package tokenizer

import (
	"fmt"
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

	tokens := []Token{}

	for _, scanner := range scanners {
		if scanner.IsToken(cleanedString) {
			tokens = append(tokens, scanner.NewToken(cleanedString))
		}
	}
	if len(tokens) > 0 {
		return tokens, nil
	}

	return nil, fmt.Errorf("no valid token %s", cleanedString)
}

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */

/**
* parents (), edged brackets <>, square brackets [], curly braces {}
 */
