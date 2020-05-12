package main

func longestCommonPrefix(strs []string) string {
	result := ""
	length := len(strs)

	if length == 0 {
		return result
	}

	if length == 1 {
		return strs[0]
	}

	rlength := len(strs[0])
	if rlength == 0 {
		return result
	}

	for j := 0; j < rlength; j++ {
		for i := 1; i < length; i++ {
			if len(strs[i]) <= j {
				return result
			}

			if strs[i][j] != strs[0][j] {
				return result
			}

			if i == length-1 {
				result = result + string(strs[0][j])
			}
		}
	}
	return result
}
