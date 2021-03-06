package ast

// OpaqueValueExpr is expression.
type OpaqueValueExpr struct {
	Addr       Address
	Pos        Position
	Type       string
	ChildNodes []Node
}

func parseOpaqueValueExpr(line string) *OpaqueValueExpr {
	groups := groupsFromRegex(
		"<(?P<position>.*)> '(?P<type>.*)'",
		line,
	)

	return &OpaqueValueExpr{
		Addr:       ParseAddress(groups["address"]),
		Pos:        NewPositionFromString(groups["position"]),
		Type:       groups["type"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *OpaqueValueExpr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *OpaqueValueExpr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *OpaqueValueExpr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *OpaqueValueExpr) Position() Position {
	return n.Pos
}
