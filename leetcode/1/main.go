package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		dif := target - num
		if j, ok := seen[dif]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return nil
}
