package p1

import "math"

func sum(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func product(nums []int) int {
	p := nums[0]
	for i := 1; i < len(nums); i++ {
		p *= nums[i]
	}
	return p
}

func mean(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total / len(nums)
}

func sqrt(num float64) float64 {
	return math.Sqrt(num)
}
