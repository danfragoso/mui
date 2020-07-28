package mui

import (
	"encoding/json"
	"fmt"
)

// Node - MUI Lang node represantation
type Node struct {
	Name     string  `json:"name"`
	Content  string  `json:"content"`
	Props    []*Prop `json:"props"`
	Children []*Node `json:"children"`
}

// NewNode -
func NewNode(name string) *Node {
	return &Node{Name: name}
}

// GetName -
func (node *Node) GetName() string {
	return node.Name
}

// SetName -
func (node *Node) SetName(name string) {
	node.Name = name
}

// GetContent -
func (node *Node) GetContent() string {
	return node.Content
}

// SetContent -
func (node *Node) SetContent(content string) {
	node.Content = content
}

// GetProps -
func (node *Node) GetProps() []*Prop {
	return node.Props
}

// GetProps -
func (node *Node) GetProp(propName string) string {
	for _, prop := range node.Props {
		if prop.Key == propName {
			return prop.Value
		}
	}

	return "nil"
}

// SetProps -
func (node *Node) SetProps(props []*Prop) {
	node.Props = props
}

// AddProp -
func (node *Node) AddProp(prop *Prop) {
	node.Props = append(node.Props, prop)
}

// AddProps -
func (node *Node) AddProps(props ...*Prop) {
	node.Props = append(node.Props, props...)
}

// GetChildren -
func (node *Node) GetChildren() []*Node {
	return node.Children
}

// SetChildren -
func (node *Node) SetChildren(children []*Node) {
	node.Children = children
}

// AddChild -
func (node *Node) AddChild(child *Node) {
	node.Children = append(node.Children, child)
}

// AddChildren -
func (node *Node) AddChildren(children ...*Node) {
	node.Children = append(node.Children, children...)
}

func (node *Node) AsJSON() string {
	res, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		fmt.Println(err)
		return "{}"
	}

	return string(res)
}
