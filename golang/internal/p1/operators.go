package p1

import "math"

func Sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func Product(nums []int) int {
	p := 1
	for _, num := range nums {
		p *= num
	}
	return p
}

func Mean(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s / len(nums)
}

func Sqrt(num float64) float64 {
	return math.Sqrt(num)
}
