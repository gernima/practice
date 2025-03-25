package combineSlice

func SumSlices(s1, s2 []int) []int {
	result := make([]int, len(s1))
	copy(result, s1)
	
	for i := 0; i < len(s2); i++ {
		result[i] = s1[i] + s2[i]
	}
	
	return result
}
