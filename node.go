package mui

// Node - MUI Lang node represantation
type Node struct {
	Name     string
	Content  string
	Props    []*Prop
	Children []*Node
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
