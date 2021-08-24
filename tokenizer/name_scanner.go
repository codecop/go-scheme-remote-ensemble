package tokenizer

import "regexp"

type NameToken struct {
	Value string
}

// TODO cleanup Tokenizer: Is*Token and New*Token can be private in all modules

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
