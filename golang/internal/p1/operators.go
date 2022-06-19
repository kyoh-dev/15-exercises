package p1

import "math"

func Sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func product(nums []int) int {
	p := nums[0]
	for i := 1; i < len(nums); i++ {
		p *= nums[i]
	}
	return p
}

func mean(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s / len(nums)
}

func sqrt(num float64) float64 {
	return math.Sqrt(num)
}
