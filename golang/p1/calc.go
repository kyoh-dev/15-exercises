package p1

func sum(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func product(nums []int) int {
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		total *= nums[i]
	}
	return total
}

func mean(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total / len(nums)
}

func sqrt(nums []int) float64 {
	return 0
}
