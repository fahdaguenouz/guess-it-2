package mathskills

import "math"


func Range(numbers []float64) (int, int) {
	x := make([]float64, len(numbers))
	y := numbers

	for i := range x {
		x[i] = float64(i)
	}

	a, b := Regression(x, y)
	next := float64(len(numbers))
	pred := a + b*next

	res := make([]float64, len(numbers))
	for i := range res {
		res[i] = y[i] - (a + b*x[i])
	}

	sd := StandardDeviation(res)

	lower := int(pred - sd)
	if lower < 1 {
		lower = 1
	}
	upper := int(pred + sd)

	return lower, upper
}

func Average(numbers []float64) float64 {
	var sum float64
	for _, value := range numbers {
		sum = value + sum
	}
	return sum / float64(len(numbers))
}

func StandardDeviation(res []float64) float64 {
	avg := Average(res)
	var sqrt float64
	for _, r := range res {
		sqrt += (r - avg) * (r - avg)
	}
	return math.Sqrt(sqrt / float64(len(res)-1))
}

func Regression(x, y []float64) (float64, float64) {
	n := len(x)

	var sumX float64
	var sumY float64
	var sumXY float64
	var sumXX float64

	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXX += x[i] * x[i]
	}

	avgX := Average(x)
	avgY := Average(y)

	beta := (sumXY - float64(n)*avgX*avgY) / (sumXX - float64(n)*avgX*avgX)
	alpha := avgY - beta*avgX

	return alpha, beta
}