package lseq

type Node struct {
	BlindNode BlindNode
	Path      []Couple
}

func NewNode(blindNode BlindNode, path []Couple) Node {
	return Node{blindNode, path}
}
