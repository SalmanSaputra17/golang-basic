package math

func Average(xs []float64) float64 {
	total := float64(0)

	for _, i := range xs {
		total += i
	}

	return total / float64(len(xs))
}

func Max(xs []float64) float64 {
	max := xs[0]

	for _, i := range xs {
		if i > max {
			max = i
		}
	}

	return max
}

func Min(xs []float64) float64 {
	min := xs[0]

	for _, i := range xs {
		if i < min {
			min = i
		}
	}

	return min
}