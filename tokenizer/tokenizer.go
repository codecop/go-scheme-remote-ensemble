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
		return []Token{NewBooleanToken(cleanedString)}, nil
	}

	if IsNumberToken(cleanedString) {
		return []Token{NewNumberToken(cleanedString)}, nil
	}

	return nil, fmt.Errorf("no valid token %s", cleanedString)
}

func NewNumberToken(token string) NumberToken {
	number, _ := strconv.Atoi(token)
	return NumberToken{Value: number}
}

func NewBooleanToken(token string) BooleanToken {
	return BooleanToken{Value: token == "#t"}
}

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */
