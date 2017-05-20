package lseq

type Couple struct {
	Pos  uint
	Site uint
}

func NewCouple(pos, site uint) Couple {
	return Couple{pos, site}
}

func MinCouple() Couple {
	return NewCouple(0, 0)
}

func MaxCouple(depth uint) Couple {
	return NewCouple(MaxSite, MaxUInt)
}

// Compare compares two Couple.
// 1 if c is greater than target
// 0 if c and target are equal
// -1 if c is smaller than target
func (c Couple) Compare(target Couple) int {
	if c.Pos > target.Pos {
		return 1
	} else if c.Pos < target.Pos {
		return -1
	} else if c.Site > target.Site {
		return 1
	} else if c.Site < target.Site {
		return -1
	}

	return 0
}
