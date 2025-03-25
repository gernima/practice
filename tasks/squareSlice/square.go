package squareSlice

func SliceNew(s []int) []int {
	result := make([]int, len(s))
	copy(result, s)

	for i := 0; i < len(s); i++ {
		result[i] = s[i] * s[i]
	}

	return result
}

func SliceInPlace(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] * s[i]
	}
}
