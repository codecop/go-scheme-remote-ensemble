package tokenizer

type ParenthesisToken struct {
	Value string
}

func IsParenthesisToken(token string) bool {
	return token == "(" || token == ")"
}

func NewParenthesisToken(token string) Token {
	parenthesis := "("
	return ParenthesisToken{Value: parenthesis}
}

func NewParenthesisScanner() Scanner {
	return Scanner{
		IsToken:  IsParenthesisToken,
		NewToken: NewParenthesisToken,
	}
}
