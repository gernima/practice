package main

// First solution
func longestCommonPrefix(strs []string) string {
	prefix := ""
	pi := ""
	isPart := false

	for i := 0; i < len(strs[0]); i++ {
		pi = string(strs[0][i])
		for j := 0; j < len(strs); j++ {
			if i < len(strs[j]) && (string(strs[j][i]) == pi || pi == "") {
				isPart = true
			} else {
				return prefix
			}
		}
		if isPart == true {
			prefix += pi
		}
	}

	return prefix
}

// Optimized solution
func newLongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs[0]) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
