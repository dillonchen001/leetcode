package main

func longestPalindrome(s string) string {
	len1 := len(s)
	if len1 <= 1 {
		return s
	}

	tmpS := "#"
	for i := 0; i < len1; i++ {
		tmpS = tmpS + string(s[i])
		tmpS = tmpS + "#"
	}

	max := 1
	mid := 0

	location := make(map[int]int)
	location[0] = 1

	resl := 1
	resi := 0

	for i := 0; i < len(tmpS)-1; i++ {
		if max > i {
			if location[2*mid-1] > max-i {
				location[i] = max - i
			} else {
				location[i] = location[2*mid-1]
			}
		} else {
			location[i] = 1
		}

		for {
			if i < location[i] || location[i]+i > len(tmpS)-1 || tmpS[i-location[i]] != tmpS[i+location[i]] {
				break
			}
			location[i]++
		}

		if location[i]+i > max {
			max = location[i] + i
			mid = i
		}

		if location[i] > resl {
			resl = location[i]
			resi = i
		}

	}

	start := (resi + 1 - resl) / 2

	return s[start : start+resl-1]
}
