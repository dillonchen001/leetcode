package main

func convert(s string, numRows int) string {
	len1 := len(s)
	if len1 <= numRows || numRows <= 1 {
		return s
	}

	stage := 2*numRows - 2
	result := ""

	for i := 0; i < numRows; i++ {
		for k := 0; k < len1; k = k + stage {
			if k+i >= len1 {
				break
			}
			result = result + string(s[k+i])

			if k+stage-i < len1 && i > 0 && i < numRows-1 {
				result += string(s[k+stage-i])
			}
		}
	}
	return result
}
