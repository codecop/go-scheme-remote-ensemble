package tokenizer

import "regexp"

type BooleanToken struct {
	Value bool
}

func IsBooleanToken(token string) bool {
	trueOrFalse := regexp.MustCompile("^#[tf]$")
	return trueOrFalse.MatchString(token)
}

func NewBooleanToken(token string) Token {
	return BooleanToken{Value: token == "#t"}
}

func NewBoolScanner() Scanner {
	return Scanner{
		IsToken:  IsBooleanToken,
		NewToken: NewBooleanToken,
	}
}
