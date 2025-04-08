package main

func pivotIndex(nums []int) int {
	var totalSum int = 0

	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}

	var leftSum int = 0

	for i := 0; i < len(nums); i++ {
		rightSum := totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
