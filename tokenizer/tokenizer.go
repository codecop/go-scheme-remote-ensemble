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
	NewNumberScanner(),
	NewBoolScanner(),
}

func Scan(cleanedString string) ([]Token, error) {
	if len(cleanedString) == 0 {
		return nil, nil
	}

	for _, scanner := range scanners {
		if scanner.IsToken(cleanedString) {
			return []Token{scanner.NewToken(cleanedString)}, nil
		}
	}

	return nil, fmt.Errorf("no valid token %s", cleanedString)
}

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */

/**
* parents (), edged brackets <>, square brackets [], curly braces {}
 */
