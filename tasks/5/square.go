package main

func sliceNew(s []int) []int {
	if s == nil {
		return nil
	}

	if len(s) == 0 {
		return []int{}
	}

	result := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		result[i] = s[i] * s[i]
	}

	return result
}

func sliceInPlace(s []int) {
	if s == nil || len(s) == 0 {
		return
	}

	for i := 0; i < len(s); i++ {
		s[i] = s[i] * s[i]
	}
}
