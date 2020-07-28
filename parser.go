package mui

import (
	"errors"
)

// Parse -
func Parse(source string) (*Node, error) {
	tokenList, err := Tokenize(source)
	var root *Node
	var currentNode *Node

	if err != nil {
		return nil, err
	}

	for tokenList.HasToken() {
		if tokenList.CurrentToken().Is(Identifier) {
			if tokenList.NextNonWhitespaceToken().
				IsExpectedAfter(tokenList.CurrentToken()) {

				if tokenList.NextNonWhitespaceToken().Is(OpenParenthesis) {
					currentNode = NewNode(tokenList.CurrentToken().Value)

					if root == nil {
						root = currentNode
					}
				} else {
					//@TODO: Create new prop
				}
			} else {
				return nil, errors.New("Unexpected token")
			}
		} else if tokenList.CurrentToken().Is(String) {
			if tokenList.PreviousNonWhitespaceToken().Is(OpenParenthesis) {
				currentNode.Content = tokenList.CurrentToken().Value
			}
		}

		tokenList.Advance()
	}

	return root, nil
}
