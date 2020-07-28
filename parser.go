package mui

import (
	"errors"
)

// Parse -
func Parse(source string) (*Node, error) {
	firstToken, err := Tokenize(source)
	if err != nil {
		return nil, err
	}

	currentToken := firstToken

	var root *Node
	var currentNode *Node

	for currentToken != nil {
		if currentToken.Is(Identifier) {
			if currentToken.NextNonWhitespaceToken().IsExpectedAfter(currentToken) {
				if currentToken.NextNonWhitespaceToken().Is(OpenParenthesis) {
					currentNode = NewNode(currentToken.Value)

					if root == nil {
						root = currentNode
					}
				} else {
					if currentToken.NextNonWhitespaceToken().NextNonWhitespaceToken().
						Is(String) {
						currentNode.AddProp(
							NewProp(
								currentToken.Value,
								currentToken.NextNonWhitespaceToken().NextNonWhitespaceToken().Value,
							),
						)
					}
				}
			} else {
				return nil, errors.New("Unexpected token")
			}
		} else if currentToken.Is(String) {
			if currentToken.PreviousNonWhitespaceToken().Is(OpenParenthesis) {
				currentNode.Content = currentToken.Value
			}
		}

		currentToken = currentToken.NextToken()
	}

	return root, nil
}
