package lseq

import (
	"errors"
	"math"
	"math/rand"
)

const MaxSite = 32767
const MaxPos = math.MaxInt32

type LSEQ struct {
	Site  uint
	Clock uint
	Step  uint
	Root  BlindNode
}

func NewLSEQ(site, clock, step uint) LSEQ {
	res := LSEQ{
		site,
		clock,
		step,
		NewBlindNode(0, 0, true),
	}

	// add the smallest and the biggest nodes
	// using BlindNode.Add directly because LSEQ.Add takes a Node so it would
	// be more complicated to init
	res.Root.Add(MinCouple(), NewBlindNode(0, 0, false))
	res.Root.Add(MaxCouple(0), NewBlindNode(MaxPos, 0, false))

	return res
}

// GetByIndex finds the nth BlindNode.
func (l LSEQ) GetByIndex(index int) (Node, bool) {
	node, index := l.Root.GetByIndex(index, make([]Couple, 0))
	return node, index == -1
}

// GetByPath finds a BlindNode in the node tree based on a path.
func (l LSEQ) GetByPath(path []Couple) (Node, bool) {
	blindNode, ok := l.Root.GetByPath(path)
	return NewNode(*blindNode, path), ok
}

func (l *LSEQ) Insert(index int, char rune) error {
	p, ok := l.GetByIndex(index)
	if !ok {
		return errors.New("could not find any Node at index" + string(index))
	}
	q, ok := l.GetByIndex(index + 1)
	if !ok {
		return errors.New("could not find any Node at index" + string(index+1))
	}

	l.Clock++
	err := l.ApplyInsert(
		NewNode(
			NewBlindNode(l.Clock, char, false),
			l.alloc(p, q),
		),
	)

	return err
}

func (l *LSEQ) ApplyInsert(node Node) error {
	if len(node.Path) == 0 {
		return errors.New("empty path")
	}

	parent, ok := l.Root.GetByPath(node.Path[:len(node.Path)-1])
	if !ok {
		return errors.New("path not found")
	}

	parent.Add(node.Path[len(node.Path)-1], node.BlindNode)
	return nil
}

func (l *LSEQ) Remove(index int) error {
	// the +1 is just a quick fix, check that we don't delete first or last
	node, ok := l.GetByIndex(index + 1)
	if !ok {
		return errors.New("path not found")
	}

	return l.ApplyRemove(node.Path)
}

func (l *LSEQ) ApplyRemove(path []Couple) error {
	parent, ok := l.Root.GetByPath(path[:len(path)-1])
	if !ok {
		return errors.New("path not found")
	}

	parent.Remove(path[len(path)-1])
	return nil
}

func (l LSEQ) alloc(p, q Node) []Couple {
	// assume p < q as Insert assures it
	maxPathLen := maxInt(len(p.Path), len(q.Path))

	for i := 0; i < maxPathLen; i++ {
		pc := getCouple(i, p.Path, false, 0)
		qc := getCouple(i, q.Path, true, uint(len(p.Path)))
		posDiff := uintDistance(pc.Pos, qc.Pos)

		if posDiff == 0 && pc.Site == qc.Site {
			// positions and sites are equal, just keep going
			continue
		} else if posDiff == 0 && pc.Site < l.Site && l.Site < qc.Site {
			// positions are equal but the site will make it go between p and q
			return changeLastCouple(p.Path, NewCouple(pc.Pos, l.Site))
		} else if posDiff == 0 {
			// it should be a child of p
			return l.allocChildOf(p)
		}

		if posDiff == 1 && pc.Site < l.Site && l.Site < MaxSite {
			// no space between p's pos and q's pos but the site will make it
			// go after p (and before q)
			return changeLastCouple(p.Path, NewCouple(pc.Pos, l.Site))
		} else if posDiff == 1 && MaxSite < l.Site && l.Site < qc.Site {
			// no space between p's pos and q's pos but the site will make it
			// go before q (and after p)
			return changeLastCouple(p.Path, NewCouple(qc.Pos, l.Site))
		} else if posDiff == 1 {
			// it should be a child of p
			return l.allocChildOf(p)
		}

		// it should be between pc.Pos and qc.Pos
		return changeLastCouple(
			p.Path,
			l.coupleBetween(pc, qc, len(p.Path)),
		)
	}

	// it means that p and q have the exact same paths
	// which would be pretty surprising...
	// should we panic?
	return []Couple{}
}

func (l LSEQ) coupleBetween(p, q Couple, currentDepth int) Couple {
	step := minInt(int(l.Step), int(uintDistance(p.Pos, q.Pos)))
	add := uint(rand.Intn(step + 1))

	var pos uint
	if chooseStrategy(currentDepth) {
		pos = p.Pos + add
	} else {
		pos = q.Pos - add
	}

	return NewCouple(pos, l.Site)
}

func (l LSEQ) allocChildOf(p Node) []Couple {
	return append(
		p.Path,
		l.coupleBetween(MinCouple(), MaxCouple(uint(len(p.Path)+1)), len(p.Path)+1),
	)
}

func (l LSEQ) ToString() string {
	return l.Root.ToString()
}
