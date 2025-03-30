package combineSlice

func sumSlices(s1, s2 []int) []int {
	if s1 == nil || s2 == nil {
		return nil
	}

	if len(s1) == 0 || len(s2) == 0 {
		return []int{}
	}

	var base, other []int
	if len(s1) >= len(s2) {
		base, other = s1, s2
	} else {
		base, other = s2, s1
	}

	result := make([]int, len(base))
	for i := range base {
		result[i] = base[i]
	}

	for i := 0; i < len(other); i++ {
		result[i] += other[i]
	}

	return result
}
