package main

func clamp(val, low, high int) int {
	switch {
	case val < low:
		return low
	case val > high:
		return high
	default:
		return val
	}
}

func between(val, low, high int) bool {
	return clamp(val, low, high) == val
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(terms ...int) int {
	result := terms[0]
	for _, term := range terms {
		if term > result {
			result = term
		}
	}
	return result
}

func min(terms ...int) int {
	result := terms[0]
	for _, term := range terms {
		if term < result {
			result = term
		}
	}
	return result
}
