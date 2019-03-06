package constants

type TokenType int

const (
	IDENTIFIER TokenType = 0
	STRING     TokenType = 1
	NUMBER     TokenType = 2

	NOT          TokenType = 3
	NOTEQUAL     TokenType = 4
	EQUAL        TokenType = 5
	DOUBLEEQUAL  TokenType = 6
	GREATER      TokenType = 7
	SMALLER      TokenType = 8
	GREATEREQUAL TokenType = 9
	SMALLEREQUAL TokenType = 10

	PLUS     TokenType = 11
	MINUS    TokenType = 12
	MULTIPLY TokenType = 13
	DIVIDE   TokenType = 14

	EOF TokenType = 15
)
