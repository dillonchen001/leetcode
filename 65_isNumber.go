package main

import "fmt"

func main() {
	s := "32.e-80123"
	fmt.Println(isNumber(s))
}

func isNumber(s string) bool {
	ls := len(s)

	symbol := false
	digit := false
	index := false
	point := 0
	for i := 0; i < ls; i++ {
		if s[i] == ' ' {
			if !digit && !symbol && point == 0 && !index {
				continue
			}

			for i := i + 1; i < ls; i++ {
				if s[i] != ' ' {
					return false
				}
			}
		} else if s[i] == '+' || s[i] == '-' {
			if point > 0 || digit {
				return false
			}

			symbol = true
		} else if s[i] == '.' {
			point++
			if point > 1 || index {
				return false
			}
		} else if s[i] == 'e' {
			if index {
				return false
			}
			if !digit {
				return false
			}
			if point > 0 {
				point--
			}
			symbol = false
			digit = false
			index = true
		} else if s[i] >= '0' && s[i] <= '9' {
			digit = true
		} else {
			return false
		}
	}
	return digit
}
