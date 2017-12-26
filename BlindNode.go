package lseq

import (
	"fmt"
	"strings"
)

// BlindNode is a part of the node tree.
// It is blind because it doesn't know the path required to get to it.
// The character is optional, a node can have as only role to contain children.
// Order is a slice of the keys from the Children map, used to order them.
type BlindNode struct {
	Char     rune
	IsRoot   bool
	Children map[Triple]*BlindNode
	Order    []Triple
}

func NewBlindNode(char rune, isRoot bool) BlindNode {
	return BlindNode{
		char,
		isRoot,
		map[Triple]*BlindNode{},
		[]Triple{},
	}
}

// Length recursively calculates the number of characters, in this node and in
// its children.
func (n BlindNode) Length() (res uint) {
	if n.IsChar() {
		res++
	}

	for _, v := range n.Children {
		res += v.Length()
	}

	return
}

// Add a children and add the Triple to n.Order.
func (n *BlindNode) Add(triple Triple, node BlindNode) {
	n.Children[triple] = &node

	// add triple in n.Order, at the right spot
	for k, v := range n.Order {
		if triple.Compare(v) == -1 {
			// insert triple in index k
			n.Order = append(n.Order, NewTriple(0, 0, 0))
			copy(n.Order[k+1:], n.Order[k:])
			n.Order[k] = triple

			return
		}
	}

	// triple is the biggest triple in n.Order
	n.Order = append(n.Order, triple)
}

func (n *BlindNode) Remove(triple Triple) {
	child, ok := n.Children[triple]
	if !ok {
		// child does not exist, no need to delete him
		return
	}

	if len(child.Children) != 0 {
		// it has children, we can't just delete it, we remove the char from
		// the node and don't remove the triple from n.Order
		child.Char = 0
		return
	}

	delete(n.Children, triple)

	var i int
	for k, v := range n.Order {
		if v == triple {
			i = k
			break
		}
	}

	n.Order = append(n.Order[:i], n.Order[i+1:]...)
}

// GetByIndex finds the nth Node.
func (n *BlindNode) GetByIndex(index int, path []Triple) (Node, int) {
	if !n.IsRoot && index == 0 {
		return NewNode(*n, path), -1
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
func (n *BlindNode) GetByPath(path []Triple) (*BlindNode, bool) {
	if len(path) == 0 {
		return n, true
	}

	nextTriple, path := path[0], path[1:]
	child, ok := n.Children[nextTriple]

	if !ok {
		return new(BlindNode), false
	}

	return child.GetByPath(path)
}

// IsChar informs us whether this BlindNode contains a character or is only a parent.
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

func (n BlindNode) Debug(triple Triple, depth int) {
	if n.IsRoot {
		depth -= 1
	} else {
		var tmp string
		if n.IsChar() {
			tmp = ": " + string(n.Char)
		}

		fmt.Printf("%v%v %v\n", strings.Repeat("\t", depth), triple, tmp)
	}

	for _, v := range n.Order {
		n.Children[v].Debug(v, depth+1)
	}
}
