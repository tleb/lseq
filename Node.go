package lseq

type Node struct {
	BlindNode BlindNode
	Path      []Triple
}

func NewNode(blindNode BlindNode, path []Triple) Node {
	return Node{blindNode, path}
}
