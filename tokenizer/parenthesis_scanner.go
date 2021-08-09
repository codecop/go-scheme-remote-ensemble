package tokenizer

type ParenthesisToken struct {
	Value string
}

func IsParenthesisToken(token string) bool {
	return token == "(" || token == ")"
}

func NewParenthesisToken(token string) Token {
	return ParenthesisToken{Value: token}
}

func NewParenthesisScanner() Scanner {
	return Scanner{
		IsToken:  IsParenthesisToken,
		NewToken: NewParenthesisToken,
	}
}
