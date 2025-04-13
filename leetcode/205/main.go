package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := make(map[byte]byte)
	s2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if val, ok := s1[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			s1[s[i]] = t[i]
		}

		if val, ok := s2[t[i]]; ok {
			if val != s[i] {
				return false
			}
		} else {
			s2[t[i]] = s[i]
		}
	}

	return true
}
