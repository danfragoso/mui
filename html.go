package mui

import "fmt"

func EmitHTML(root *Node) (string, error) {
	return generateHTML(root)
}

func generateHTML(node *Node) (string, error) {
	fmt.Println(node)
	fmt.Println(generateNode(node))
	return "", nil
}

func generateNode(node *Node) string {
	var chilrenNodes string
	for _, child := range node.GetChildren() {
		chilrenNodes += generateNode(child)
	}

	return `<` + node.Name + generateProps(node) + `>` + node.Content + `
` + chilrenNodes + `
</` + node.Name + `>`
}

func generateProps(node *Node) string {
	var propString string
	props := node.GetProps()

	for _, prop := range props {
		propString += " " + prop.Key + "=" + "\"" + prop.Value + "\"" + " "
	}

	return propString
}
