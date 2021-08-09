package tokenizer

import "strconv"

type NumberToken struct {
	Value int
}

func IsNumberToken(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func NewNumberToken(token string) Token {
	number, _ := strconv.Atoi(token)
	return NumberToken{Value: number}
}

func NewNumberScanner() Scanner {
	return Scanner{
		IsToken:  IsNumberToken,
		NewToken: NewNumberToken,
	}
}
