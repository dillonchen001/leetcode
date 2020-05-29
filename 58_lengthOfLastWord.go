package main

func lengthOfLastWord(s string) int {
	ret := 0
	ls := len(s)
	if ls == 0 {
		return ret
	}

	for i := ls - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ret > 0 {
				return ret
			}
		} else {
			ret++
		}
	}
	return ret
}
