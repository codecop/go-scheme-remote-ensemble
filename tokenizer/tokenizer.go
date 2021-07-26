package tokenizer

import (
	"fmt"
	"strconv"
)

type Token interface {
}

func IsNumberToken(cleanedString string) bool {
	_, err := strconv.Atoi(cleanedString)
	return err == nil
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
	if cleanedString == "#t" {
		return []Token{BooleanToken{Value: true}}, nil
	}
	if cleanedString == "#f" {
		return []Token{BooleanToken{Value: false}}, nil
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
