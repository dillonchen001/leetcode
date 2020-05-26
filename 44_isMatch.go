package main

func isMatch(s string, p string) bool {
	lp := len(p)
	ls := len(s)

	start_p := -1
	mark_s := 0

	i := 0
	j := 0
	for i < ls {
		if j >= lp {
			if start_p != -1 {
				i = mark_s + 1
				j = start_p + 1
				mark_s++
				continue
			} else {
				return false
			}
		}

		if s[i] == p[j] || p[j] == '?' {
			i++
			j++
		} else if p[j] == '*' {
			start_p = j
			mark_s = i
			j++
		} else if start_p != -1 {
			i = mark_s + 1
			j = start_p + 1
			mark_s++
		} else {
			return false
		}
	}

	if j < lp {
		if p[j] == '*' {
			j++
		} else {
			return false
		}
	}

	return true
}
