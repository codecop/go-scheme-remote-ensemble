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

	//TODO refactor to align isFuncs and newTokenFuncs
	//TODO create slice of structs of functions

	for i, isFunc := range isFuncs {
		if isFunc(cleanedString) {
			return []Token{newTokenFuncs[i](cleanedString)}, nil
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

/**
* TODO: create slice in the beginning and append afterwards with tokens
 */

/**
* parents (), edged brackets <>, square brackets [], curly braces {}
 */
