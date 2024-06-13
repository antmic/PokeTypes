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
