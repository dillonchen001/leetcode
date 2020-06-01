package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	srMap := make(map[byte]int)
	deMap := make(map[byte]int)

	lt := len(t)

	for i := 0; i < lt; i++ {
		srMap[t[i]]++
	}

	minn := -1
	num := 0
	k := 0
	left := 0
	right := 0
	for j := 0; j < len(s); j++ {
		deMap[s[j]]++

		if deMap[s[j]] <= srMap[s[j]] {
			num++
		}

		if num == lt {
			for k <= j && deMap[s[k]] > srMap[s[k]] {
				deMap[s[k]]--
				k++
			}

			if minn == -1 || minn > j-k+1 {
				minn = j - k + 1
				left = k
				right = j
			}

			for k <= j && num == lt {
				deMap[s[k]]--

				if deMap[s[k]] < srMap[s[k]] {
					num--
					k++
				}
			}
		}
	}

	if minn != -1 {
		return s[left : right+1]
	}
	return ""
}
