package mui

// Node - MUI Lang node represantation
type Node struct {
	name     string
	content  string
	props    []*Prop
	children []*Node
}

// NewNode -
func NewNode(name string) *Node {
	return &Node{name: name}
}

// GetName -
func (node *Node) GetName() string {
	return node.name
}

// SetName -
func (node *Node) SetName(name string) {
	node.name = name
}

// GetContent -
func (node *Node) GetContent() string {
	return node.content
}

// SetContent -
func (node *Node) SetContent(content string) {
	node.content = content
}

// GetProps -
func (node *Node) GetProps() []*Prop {
	return node.props
}

// SetProps -
func (node *Node) SetProps(props []*Prop) {
	node.props = props
}

// AddProp -
func (node *Node) AddProp(prop *Prop) {
	node.props = append(node.props, prop)
}

// AddProps -
func (node *Node) AddProps(props ...*Prop) {
	node.props = append(node.props, props...)
}

// GetChildren -
func (node *Node) GetChildren() []*Node {
	return node.children
}

// SetChildren -
func (node *Node) SetChildren(children []*Node) {
	node.children = children
}

// AddChild -
func (node *Node) AddChild(child *Node) {
	node.children = append(node.children, child)
}

// AddChildren -
func (node *Node) AddChildren(children ...*Node) {
	node.children = append(node.children, children...)
}
