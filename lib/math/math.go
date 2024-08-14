package math

import (
	"math"
	"sort"
)

// Number Constraints for floating-point and integer types
type Number interface {
	int | float64
}

// Mean Generic function to calculate the mean
func Mean[T Number](data []T) float64 {
	sum := T(0)
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}

// Median Generic function to calculate the median
func Median[T Number](data []T) float64 {
	n := len(data)
	if n == 0 {
		return 0.0
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	if n%2 == 0 {
		return float64(data[n/2-1]+data[n/2]) / 2.0
	}
	return float64(data[n/2])
}

// Max Generic function to return the higher of two numbers
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min Generic function to return the higher of two numbers
func Min[T Number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func Float32ToUint32(f float32) uint32 {
	return math.Float32bits(f)
}

func Uint32ToFloat32(u uint32) float32 {
	return math.Float32frombits(u)
}
