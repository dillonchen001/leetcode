package main

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	last := countAndSay(n - 1)
	ln := len(last)

	result := ""
	for i := 0; i < ln-1; i++ {
		for j := i + 1; j < ln; j++ {
			if last[i] == last[j] {
				if j == ln-1 {
					result += strconv.Itoa(j-i+1) + string(last[i])
					return result
				}
				continue
			} else {
				result += strconv.Itoa(j-i) + string(last[i])
				if j == ln-1 {
					result += strconv.Itoa(1) + string(last[j])
					return result
				}
				i = j - 1
				break
			}
		}
	}
	return result
}
