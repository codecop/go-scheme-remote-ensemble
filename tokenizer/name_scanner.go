package tokenizer

import "regexp"

type NameToken struct {
	Value string
}

func IsNameToken(token string) bool {
	identifier := regexp.MustCompile("^[a-zA-Z]+$")
	return identifier.MatchString(token)
}

func NewNameToken(token string) Token {
	return NameToken{Value: token}
}

func NewNameScanner() Scanner {
	return Scanner{
		IsToken:  IsNameToken,
		NewToken: NewNameToken,
	}
}
