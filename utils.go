package lseq

func minInt(ints ...int) int {
	// assume len(ints) > 0
	min := ints[0]

	for _, v := range ints {
		if v < min {
			min = v
		}
	}

	return min
}

func maxInt(ints ...int) int {
	// assume len(ints) > 0
	max := ints[0]

	for _, v := range ints {
		if v > max {
			max = v
		}
	}

	return max
}

func chooseStrategy(depth int) bool {
	return depth%2 == 0
}

// getCouple gets path[i] or, if it doesn't exist, get the smallest/biggest
// couple possible (max defines smallest or biggest).
func getCouple(index int, path []Couple, max bool, depth uint) Couple {
	if len(path) > index {
		return path[index]
	}

	if max {
		return MaxCouple(depth)
	}

	return MinCouple()
}

func uintDistance(a, b uint) uint {
	if a > b {
		return a - b
	}

	return b - a
}

func intDistance(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func changeLastCouple(path []Couple, couple Couple) []Couple {
	return append(path[:len(path)-1], couple)
}
