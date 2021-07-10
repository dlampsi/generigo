package generigo

// IsOdd Check if number is odd.
func IsOdd(number int) bool {
	return !(number%2 == 0)
}

// AverageFloat64 Calculate average value in list of float64.
func AverageFloat64(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
