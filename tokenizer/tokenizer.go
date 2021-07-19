package tokenizer

import "strconv"

type Token interface {
}

type NumberToken struct {
	Value int
}

func Scan(cleanedString string) ([]Token, error) {
	if len(cleanedString) == 0 {
		return nil, nil
	}

	number, _ := strconv.Atoi(cleanedString)
	return []Token{NumberToken{Value: number}}, nil
}
