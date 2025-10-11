package main

func sum(nums ...int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}

	return total
}
