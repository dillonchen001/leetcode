package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)

	if l3 != l1+l2 {
		return false
	}
	if l1 == 0 {
		return s2 == s3
	}

	if l2 == 0 {
		return s1 == s3
	}

	if s1[l1-1] == s3[l3-1] && isInterleave(s1[0:l1-1], s2, s3[0:l3-1]) {
		return true
	}

	if s2[l2-1] == s3[l3-1] && isInterleave(s1, s2[0:l2-1], s3[0:l3-1]) {
		return true
	}

	return false
}
