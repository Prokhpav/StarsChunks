package run

func Max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func Cmp(a, min, max int) int {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}
