package mui

import "fmt"

// Tokenize -
func Tokenize(source string) ([]*Token, error) {
	fmt.Println(source)
	fmt.Print(tokenRegex[Name].FindAllString(source, -1))
	return nil, nil
}
