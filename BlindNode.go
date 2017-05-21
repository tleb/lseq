package lseq

// BlindNode is a part of the node tree.
// It is blind because it doesn't know the path required to get to it.
// The character is optional, a node can have as only role to contain children.
// Order is a slice of the keys from the Children map, used to order them.
type BlindNode struct {
	Clock    uint
	Char     rune
	IsRoot   bool
	Children map[Couple]BlindNode
	Order    []Couple
}

func NewBlindNode(clock uint, char rune, isRoot bool) BlindNode {
	return BlindNode{
		clock,
		char,
		isRoot,
		map[Couple]BlindNode{},
		[]Couple{},
	}
}

// Length recursively calculates the number of characters, in this node and in
// its children.
func (n BlindNode) Length() uint {
	var res uint

	if n.IsChar() {
		res++
	}

	for _, v := range n.Children {
		res += v.Length()
	}

	return res
}

// Add a children and adds the Couple to n.Order.
func (n *BlindNode) Add(couple Couple, node BlindNode) {
	n.Children[couple] = node

	// add couple in n.Order, at the right spot
	for k, v := range n.Order {
		if couple.Compare(v) == -1 {
			// insert couple in index k
			n.Order = append(n.Order, NewCouple(0, 0))
			copy(n.Order[k+1:], n.Order[k:])
			n.Order[k] = couple

			return
		}
	}

	// couple is the biggest couple in n.Order
	n.Order = append(n.Order, couple)
}

func (n *BlindNode) Remove(couple Couple) {
	child, ok := n.Children[couple]
	if !ok {
		// child does not exist, no need to delete him
		return
	}

	if len(child.Children) != 0 {
		// it has children, we can't just delete it, we remove the char from
		// the node and don't remove the couple from n.Order
		child.Char = 0
		return
	}

	delete(n.Children, couple)

	var i int
	for k, v := range n.Order {
		if v == couple {
			i = k
			break
		}
	}

	n.Order = append(n.Order[:i], n.Order[i+1:]...)
}

// GetByIndex finds the nth Node.
func (n BlindNode) GetByIndex(index int, path []Couple) (Node, int) {
	if !n.IsRoot && index == 0 {
		return NewNode(n, path), -1
	}

	if !n.IsRoot {
		index--
	}

	var res Node
	for _, v := range n.Order {
		res, index = n.Children[v].GetByIndex(index, append(path, v))

		if index == -1 {
			return res, -1
		}
	}

	return Node{}, index
}

// GetByPath finds a BlindNode in the node tree based on a path.
func (n *BlindNode) GetByPath(path []Couple) (*BlindNode, bool) {
	if len(path) == 0 {
		return n, true
	}

	nextCouple, path := path[0], path[1:]
	child, ok := n.Children[nextCouple]

	if !ok {
		return new(BlindNode), false
	}

	return child.GetByPath(path)
}

// IsChar informs us whether this BlindNode contains a character or only is parent.
func (n BlindNode) IsChar() bool {
	return !n.IsRoot && n.Char != 0
}

func (n BlindNode) ToString() string {
	var res string

	if n.IsChar() {
		res += string(n.Char)
	}

	for _, v := range n.Order {
		res += n.Children[v].ToString()
	}

	return res
}
