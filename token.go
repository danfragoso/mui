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
}

func (token *Token) Is(tokenType TokenType) bool {
	return token.TokenType == tokenType
}

func (token *Token) IsWhiteSpace() bool {
	return token.TokenType == NewLine || token.TokenType == Tab || token.TokenType == Space
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

func (tokenList *TokenList) PreviousToken() *Token {
	return tokenList.Tokens[tokenList.tokenIdx-1]
}

func (tokenList *TokenList) NextToken() *Token {
	return tokenList.Tokens[tokenList.tokenIdx+1]
}

func (tokenList *TokenList) PreviousNonWhitespaceToken() *Token {
	for idx := tokenList.tokenIdx - 1; idx < tokenList.tokenListLen; idx-- {
		if !tokenList.Tokens[idx].IsWhiteSpace() {
			return tokenList.Tokens[idx]
		}
	}

	return nil
}

func (tokenList *TokenList) NextNonWhitespaceToken() *Token {
	for idx := tokenList.tokenIdx + 1; idx < tokenList.tokenListLen; idx++ {
		if !tokenList.Tokens[idx].IsWhiteSpace() {
			return tokenList.Tokens[idx]
		}
	}

	return nil
}

func (tokenList *TokenList) CurrentToken() *Token {
	return tokenList.Tokens[tokenList.tokenIdx]
}

func (tokenList *TokenList) Advance() {
	tokenList.tokenIdx++
}

func (tokenList *TokenList) AdvanceUntilNonWhitespace() {
	tokenList.tokenIdx++

	for !tokenList.Tokens[tokenList.tokenIdx].IsWhiteSpace() {
		tokenList.tokenIdx++
	}
}

func (tokenList *TokenList) AdvanceUntilLastWhitespace() {
	tokenList.tokenIdx++

	for !tokenList.Tokens[tokenList.tokenIdx].IsWhiteSpace() {
		tokenList.tokenIdx++
	}

	tokenList.tokenIdx--
}
