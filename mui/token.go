package mui

import "regexp"

type TokenType int

var tokenRegex map[TokenType]*regexp.Regexp = map[TokenType]*regexp.Regexp{
	Name: regexp.MustCompile("[a-zA-z]+\\d*(\\s*\\()"),
}

const (
	Name = iota
	OpenParanthesis
	CloseParenthesis
	String
	Property
	NewLine
	Comma
)

type Token struct {
	tokenType   TokenType
	value       string
	absoluteIdx int
	lineIdx     int
}
