package tokenizer

import "regexp"

type BooleanToken struct {
	Value bool
}

func isBooleanToken(token string) bool {
	trueOrFalse := regexp.MustCompile("^#[tf]$")
	return trueOrFalse.MatchString(token)
}

func newBooleanToken(token string) Token {
	return BooleanToken{Value: token == "#t"}
}

// TODO cleanup Tokenizer: Inconsistent name NewBoolScanner <-> NewBooleanScanner

func NewBoolScanner() Scanner {
	return Scanner{
		IsToken:  isBooleanToken,
		NewToken: newBooleanToken,
	}
}
