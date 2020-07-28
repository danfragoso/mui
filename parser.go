package mui

// Parse -
func Parse(source string) (*Node, error) {
	tokenList, err := Tokenize(source)
	var root *Node
	//var currentNode *Node

	if err != nil {
		return nil, err
	}

	for tokenList.HasToken() {

		tokenList.Advance()
	}

	return root, nil
}
