package utils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Avg[T Number](nums []T) float64 {
	var sum T
	for _, v := range nums {
		sum += v
	}
	return float64(sum) / float64(len(nums))
}
