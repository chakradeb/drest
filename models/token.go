package models

import "github.com/chakradeb/drest/constants"

type Token struct {
	TokenType  constants.TokenType
	Lexeme     string
	LineNumber int
}

func NewToken(t constants.TokenType, l string, line int) *Token {
	return &Token{
		TokenType: t,
		Lexeme: l,
		LineNumber: line,
	}
}
