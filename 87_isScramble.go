package main

func isScramble(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)

	if l1 != l2 {
		return false
	}

	if s1 == s2 {
		return true
	}

	if strSort(s1) != strSort(s2) {
		return false
	}

	for i := 1; i <= li-1; i++ {
		s11 := s1[:i]
		s12 := s1[i:]

		s21 := s2[:i]
		s22 := s2[i:]

		if isScramble(s11, s21) && isScramble(s12, s22) {
			return true
		}

		s21 := s2[l2-i:]
		s22 := s2[:l2-i]

		if isScramble(s11, s21) && isScramble(s12, s22) {
			return true
		}
	}
	return false

}

func strSort(s string) string {
	ls := len(s)
	if ls <= 1 {
		return s
	}

	var strl []string
	for i := 0; i < ls; i++ {
		strl = append(strl, string(s[i]))
	}

	sort.Strings(strl)

	var ret string
	for _, val := range strl {
		ret += val
	}
	return ret
}
