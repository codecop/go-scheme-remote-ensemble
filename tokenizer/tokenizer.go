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

func Scan(cleanedString string) ([]Token, error) {
	if len(cleanedString) == 0 {
		return nil, nil
	}

	scanners := []Scanner{NewNumberScanner(), NewBoolScanner()}

	//TODO refactor to align isFuncs and newTokenFuncs
	//TODO create slice of structs of functions

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
