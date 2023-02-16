package sfolib

// abs func only for int values
func Abs(n int) int {
	switch {
	case n == 0:
		return 0
	case n > 0:
		return n
	case n < 0:
		return n - n*2
	}
	panic("unreacheble")
}

// find percent between to values
func Perc(n float64, a float64) float64 {
	return n / a * 100
}
