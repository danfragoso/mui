package mui

import (
	"encoding/json"
	"fmt"
)

type TokenType int

const EOF = rune(0)

const (
	Identifier = iota
	OpenParenthesis
	CloseParenthesis
	String
	Colon
	NewLine
	Tab
	Space
	Comma
)

type Token struct {
	TokenType TokenType
	Value     string
	Idx       int

	nextToken     *Token
	previousToken *Token
}

func (token *Token) Is(tokenType TokenType) bool {
	return token.TokenType == tokenType
}

func (token *Token) IsWhiteSpace() bool {
	return token.TokenType == NewLine || token.TokenType == Tab || token.TokenType == Space
}

func (token *Token) HasNextToken() bool {
	return token.nextToken != nil
}

func (token *Token) HasPreviousToken() bool {
	return token.previousToken != nil
}

func (token *Token) NextNonWhitespaceToken() *Token {
	cToken := token.nextToken

	for cToken.HasNextToken() {
		if !cToken.IsWhiteSpace() {
			return cToken
		}

		cToken = cToken.nextToken
	}

	return nil
}

func (token *Token) PreviousNonWhitespaceToken() *Token {
	cToken := token.previousToken

	for cToken.HasPreviousToken() {
		if !cToken.IsWhiteSpace() {
			return cToken
		}

		cToken = cToken.previousToken
	}

	return nil
}

func (token *Token) NextToken() *Token {
	return token.nextToken
}

func (token *Token) IsExpectedAfter(expected *Token) bool {
	switch expected.TokenType {
	case Identifier:
		if token.TokenType == Colon || token.TokenType == OpenParenthesis {
			return true
		}

		break
	}

	return false
}

func (token *Token) AsJSON() string {
	res, err := json.Marshal(token)
	if err != nil {
		fmt.Println(err)
		return "{}"
	}

	return string(res)
}

type TokenList struct {
	Tokens       []*Token
	tokenIdx     int
	tokenListLen int
}

func CreateTokenList() *TokenList {
	var tokens []*Token
	return &TokenList{Tokens: tokens}
}

func (tokenList *TokenList) Push(token *Token) {
	tokenList.Tokens = append(tokenList.Tokens, token)
	tokenList.tokenListLen = len(tokenList.Tokens)
}

func (tokenList *TokenList) HasToken() bool {
	if tokenList.tokenIdx == tokenList.tokenListLen {
		return false
	}

	return true
}
