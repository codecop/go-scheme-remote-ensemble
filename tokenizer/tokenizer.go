package tokenizer

type Token interface {
}

type NumberToken struct {
	Value int
}

func Scan(cleanedString string) ([]Token, error) {
	if len(cleanedString) == 0 {
		return nil, nil
	}
	return []Token{NumberToken{Value: 1}}, nil
}
