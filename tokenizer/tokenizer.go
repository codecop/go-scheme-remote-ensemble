package tokenizer

import (
	"fmt"
	"regexp"
	"strconv"
)

type Token interface {
}

func IsNumberToken(cleanedString string) bool {
	_, err := strconv.Atoi(cleanedString)
	return err == nil
}

func IsBooleanToken(cleanedString string) bool {
	reg := regexp.MustCompile("(#t|#f)")
	return reg.MatchString(cleanedString)
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
