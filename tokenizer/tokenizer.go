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

	isFuncs := []func(token string) bool{
		IsBooleanToken,
		IsNumberToken,
	}

	newTokenFuncs := []func(token string) Token{
		NewBooleanToken,
		NewNumberToken,
	}

	if isFuncs[0](cleanedString) {
		return []Token{newTokenFuncs[0](cleanedString)}, nil
	}

	if isFuncs[1](cleanedString) {
		return []Token{newTokenFuncs[1](cleanedString)}, nil
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

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */

/**
* parents (), edged brackets <>, square brackets [], curly braces {}
 */
