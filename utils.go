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
