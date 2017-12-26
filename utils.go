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
	// == 1 so that we start with boundary+
	return depth%2 == 1
}

// getTriple gets path[i] or, if it doesn't exist, get the smallest/biggest
// triple possible (max defines smallest or biggest).
func getTriple(index int, path []Triple, max bool, depth uint) Triple {
	if len(path) > index {
		return path[index]
	}

	if max {
		return MaxTriple(depth)
	}

	return MinTriple()
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

func changeLastTriple(path []Triple, triple Triple) []Triple {
	return append(path[:len(path)-1], triple)
}
