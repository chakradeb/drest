package models

import (
	"fmt"
	"github.com/chakradeb/drest/constants"
)

type Scanner struct {
	Source string
	Tokens []*Token
	start int
	current int
	line int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source: source,
		Tokens: nil,
		start: 0,
		current: 0,
		line: 1,
	}
}

func (s Scanner) ScanTokens() []*Token {
	for !s.isEnd() {
		s.start = s.current
		s.scanToken();
	}
	tokens := append(s.Tokens, NewToken(constants.EOF, "", s.line ))
	return tokens
}

func (s Scanner) scanToken() {
	switch char := s.advance(); char {
	case "+":
		s.addToken(constants.PLUS)
	case "-":
		s.addToken(constants.MINUS)
	case "*":
		s.addToken(constants.MULTIPLY)
	case "!":
		s.addToken(s.getType(char, constants.NOTEQUAL, constants.NOT))
	case "=":
		s.addToken(s.getType(char, constants.DOUBLEEQUAL, constants.EQUAL))
	case "<":
		s.addToken(s.getType(char, constants.SMALLEREQUAL, constants.SMALLER))
	case ">":
		s.addToken(s.getType(char, constants.GREATEREQUAL, constants.GREATER))
	case "\n":
		s.line++
	default:
		fmt.Println("Unexpected Error")
	}
}

func (s Scanner) addToken(tokenType constants.TokenType) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, NewToken(tokenType, text, s.line))
}

func (s Scanner) advance() string {
	s.current++
	return string(s.Source[s.current])
}

func (s Scanner) isEnd() bool {
	return s.current > len(s.Source)
}

func (s Scanner) match(expected string) bool {
	if s.isEnd() {
		return false
	}
	if string(s.Source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s Scanner) getType(str string, typeOne constants.TokenType, typeTwo constants.TokenType) constants.TokenType {
	var tkn constants.TokenType
	if s.match("=") {
		tkn = typeOne
	} else {
		tkn = typeTwo
	}
	return tkn
}