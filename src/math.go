package main

import (
	"math"
)

func round2(x float64) float64 {
	return math.Round(x*100) / 100
}

func sumSlice(slice []float64) float64 {
	total := 0.0
	for _, value := range slice {
		total += value
	}
	return total
}

// returns the maximum value in a slice
func maxSlice(slice []float64) float64 {
	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}
