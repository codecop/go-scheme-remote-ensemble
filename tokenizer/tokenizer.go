package tokenizer

import (
	"fmt"
	"regexp"
	"strconv"
)

type Token interface {
}

func IsNumberToken(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func IsBooleanToken(token string) bool {
	trueOrFalse := regexp.MustCompile("^#[tf]$")
	return trueOrFalse.MatchString(token)
}

type NumberToken struct {
	Value int
}

type BooleanToken struct {
	Value bool
}

type Scanner struct {
	IsToken  func(token string) bool
	NewToken func(token string) Token
}

func NewBoolScanner() Scanner {
	return Scanner{
		IsToken:  IsBooleanToken,
		NewToken: NewBooleanToken,
	}
}

func NewNumberScanner() Scanner {
	return Scanner{
		IsToken:  IsNumberToken,
		NewToken: NewNumberToken,
	}
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

func NewNumberToken(token string) Token {
	number, _ := strconv.Atoi(token)
	return NumberToken{Value: number}
}

func NewBooleanToken(token string) Token {
	return BooleanToken{Value: token == "#t"}
}

// name token:

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */

/**
* parents (), edged brackets <>, square brackets [], curly braces {}
 */
