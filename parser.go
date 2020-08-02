package mui

import (
	"errors"
)

//ParserState - Holds the current state of the parser across all parsing contexts
type ParserState struct {
	rootNode    *Node
	currentNode *Node

	currentToken *Token
}

// Parse -
func Parse(source string) (*Node, error) {
	var e error

	s := &ParserState{}
	s.currentToken, e = Tokenize(source)
	if e != nil {
		return nil, e
	}

	for s.currentToken != nil {
		var n *Node

		switch s.currentToken.TokenType {
		case Identifier:
			n, e = parseIdentifier(s)
		case String:
			n, e = parseString(s)
		case OpenParenthesis:
			n, e = parseOpenParenthesis(s)
		case CloseParenthesis:
			n, e = parseCloseParenthesis(s)
		}

		if parserShouldReturn(n, e) {
			return n, e
		}

		s.currentToken = s.currentToken.NextToken()
	}

	return s.rootNode, nil
}

func parserShouldReturn(n *Node, e error) bool {
	return n != nil || e != nil
}

func parseCloseParenthesis(s *ParserState) (*Node, error) {
	s.currentNode = s.currentNode.Parent

	return nil, nil
}

func parseOpenParenthesis(s *ParserState) (*Node, error) {
	if n, e := catchOpenParenthesisErrors(s); parserShouldReturn(n, e) {
		return n, e
	}

	return nil, nil
}

func parseString(s *ParserState) (*Node, error) {
	if n, e := catchStringErrors(s); parserShouldReturn(n, e) {
		return n, e
	}

	if s.currentToken.PreviousNonWhitespaceToken().Is(OpenParenthesis) {
		s.currentNode.Content = s.currentToken.Value
	}

	return nil, nil
}

func parseIdentifier(s *ParserState) (*Node, error) {
	if n, e := catchIdentifierErrors(s); parserShouldReturn(n, e) {
		return n, e
	}

	if s.currentToken.NextNonWhitespaceToken().Is(OpenParenthesis) {
		createdNode := NewNode(s.currentToken.Value)

		if s.currentNode != nil {
			s.currentNode.AddChild(createdNode)
		}

		s.currentNode = createdNode

		if s.rootNode == nil {
			s.rootNode = s.currentNode
		}
	} else {
		if s.currentToken.NextNonWhitespaceToken().NextNonWhitespaceToken().Is(String) {
			s.currentNode.AddProp(
				NewProp(
					s.currentToken.Value,
					s.currentToken.NextNonWhitespaceToken().NextNonWhitespaceToken().Value,
				),
			)
		}
	}

	return nil, nil
}

func catchIdentifierErrors(s *ParserState) (*Node, error) {
	if s.currentToken.NextNonWhitespaceToken() == nil {
		return nil, missingAfterError("token", "identifier", "')' or ':'", s.currentToken.Idx)
	}

	if !s.currentToken.NextNonWhitespaceToken().IsExpectedAfter(s.currentToken) {
		return nil, errors.New("Unexpected token")
	}

	return nil, nil
}

func catchStringErrors(s *ParserState) (*Node, error) {
	if s.currentToken.NextNonWhitespaceToken() == nil {
		return nil, errors.New("Unexpected end of input")
	}

	return nil, nil
}

func catchOpenParenthesisErrors(s *ParserState) (*Node, error) {
	if s.currentToken.NextNonWhitespaceToken() == nil {
		return nil, errors.New("Unexpected end of input")
	}

	return nil, nil
}
