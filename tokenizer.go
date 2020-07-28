package mui

// Tokenize -
func Tokenize(raw string) (*Token, error) {
	source := createSource(raw)
	tokenList := CreateTokenList()

	var openString bool
	var stringBuffer []rune
	var identifierBuffer []rune

	for source.hasChar() {
		currentChar := source.getChar()

		if openString {
			if isStringDelimiter(currentChar) {
				openString = false
				tokenList.Push(createStringToken(stringBuffer, source.currentIdx))
				stringBuffer = nil
			} else {
				stringBuffer = append(stringBuffer, currentChar)
			}
		} else {
			if isLetter(currentChar) {
				identifierBuffer = append(identifierBuffer, currentChar)
			} else {
				if len(identifierBuffer) > 0 {
					tokenList.Push(createIdentifierToken(identifierBuffer, source.currentIdx))
					identifierBuffer = nil
				}

				if isWhitespace(currentChar) {
					tokenList.Push(createWhitespaceToken(currentChar, source.currentIdx))
				} else if isPunct(currentChar) {
					tokenList.Push(createPunctToken(currentChar, source.currentIdx))
				} else if isStringDelimiter(currentChar) {
					if openString {
						openString = false
					} else {
						openString = true
					}
				}
			}
		}
	}

	return chainTokens(tokenList), nil
}

func chainTokens(tokenList *TokenList) *Token {
	var fstToken *Token

	for idx, token := range tokenList.Tokens {
		if idx == 0 {
			fstToken = token
		} else {
			token.previousToken = tokenList.Tokens[idx-1]
		}

		if idx+1 < len(tokenList.Tokens) {
			token.nextToken = tokenList.Tokens[idx+1]
		}
	}

	return fstToken
}

func createWhitespaceToken(ch rune, idx int) *Token {
	token := &Token{
		Value: string(ch),
		Idx:   idx,
	}

	switch ch {
	case ' ':
		token.TokenType = Space
	case '\t':
		token.TokenType = Tab
	case '\n':
		token.TokenType = NewLine
	}

	return token
}

func createPunctToken(ch rune, idx int) *Token {
	token := &Token{
		Value: string(ch),
		Idx:   idx,
	}

	switch ch {
	case ':':
		token.TokenType = Colon
	case '(':
		token.TokenType = OpenParenthesis
	case ')':
		token.TokenType = CloseParenthesis
	case ',':
		token.TokenType = Comma
	}

	return token
}

func createIdentifierToken(identifierBuffer []rune, idx int) *Token {
	return &Token{
		Value:     string(identifierBuffer),
		TokenType: Identifier,
		Idx:       idx,
	}
}

func createStringToken(stringBuffer []rune, idx int) *Token {
	return &Token{
		Value:     string(stringBuffer),
		TokenType: String,
		Idx:       idx,
	}
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isPunct(ch rune) bool {
	return ch == ':' || ch == '(' || ch == ')' || ch == ','
}

func isStringDelimiter(ch rune) bool {
	return ch == '\''
}
