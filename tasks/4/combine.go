package main

func sumSlices(s1, s2 []int) []int {
	if s1 == nil && s2 == nil {
		return nil
	}

	if s1 == nil || len(s1) == 0 {
		result := make([]int, len(s2))
		copy(result, s2)
		return result
	}

	if s2 == nil || len(s2) == 0 {
		result := make([]int, len(s1))
		copy(result, s1)
		return result
	}

	maxLen := len(s1)
	if len(s2) > maxLen {
		maxLen = len(s2)
	}

	result := make([]int, maxLen)

	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}

	for i := 0; i < minLen; i++ {
		result[i] = s1[i] + s2[i]
	}

	for i := minLen; i < maxLen; i++ {
		if len(s1) > len(s2) {
			result[i] = s1[i]
		} else {
			result[i] = s2[i]
		}
	}

	return result
}
