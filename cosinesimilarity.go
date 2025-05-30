package vector

import (
	"math"
)

func dotProduct(a, b []float32) float64 {
	var sum float64
	for i := range a {
		sum += float64(a[i] * b[i])
	}
	return sum
}

func magnitude(a []float32) float64 {
	var sum float64
	for i := range a {
		sum += float64(a[i] * a[i])
	}
	return math.Sqrt(sum)
}

func CosineSimilarity(a, b []float32) float64 {
	return dotProduct(a, b) / (magnitude(a) * magnitude(b))
}
