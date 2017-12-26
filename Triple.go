package lseq

type Triple struct {
	Pos     uint
	Site    uint
	Counter uint
}

func NewTriple(pos, site, counter uint) Triple {
	return Triple{pos, site, counter}
}

func MinTriple() Triple {
	return NewTriple(0, 0, 0)
}

func MaxTriple(depth uint) Triple {
	// TODO: use depth to determine
	return NewTriple(MaxPos, MaxSite, MaxPos)
}

// Compare compares two Triple.
// 1 if c is greater than target
// 0 if c and target are equal
// -1 if c is smaller than target
func (c Triple) Compare(target Triple) int {
	if c.Pos > target.Pos {
		return 1
	} else if c.Pos < target.Pos {
		return -1
	} else if c.Site > target.Site {
		return 1
	} else if c.Site < target.Site {
		return -1
	} else if c.Counter > target.Counter {
		return 1
	} else if c.Counter < target.Counter {
		return -1
	}

	return 0
}
