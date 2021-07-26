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

func Scan(cleanedString string) ([]Token, error) {
	if len(cleanedString) == 0 {
		return nil, nil
	}
	// TODO: extract token classification

	if IsBooleanToken(cleanedString) {
		return []Token{BooleanToken{Value: cleanedString == "#t"}}, nil
	}

	if IsNumberToken(cleanedString) {
		number, _ := strconv.Atoi(cleanedString)
		return []Token{NumberToken{Value: number}}, nil
	}

	return nil, fmt.Errorf("no valid token %s", cleanedString)

}

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */
