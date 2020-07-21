package mui

// Parse -
func Parse(source string) (*Node, error) {
	Tokenize(source)
	return NewNode("memes"), nil
}
