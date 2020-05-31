package main

import "fmt"

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)

	if la > lb {
		return addBinary(b, a)
	}

	cur := 0
	ret := ""
	for i := la - 1; i >= 0; i-- {
		if a[i] == '1' {
			cur++
		}

		if b[lb-1+i-la+1] == '1' {
			cur++
		}

		if cur == 0 || cur == 2 {
			ret = "0" + ret
		} else {
			ret = "1" + ret
		}

		if cur >= 2 {
			cur = 1
		} else {
			cur = 0
		}
	}

	for k := lb - la - 1; k >= 0; k-- {
		if b[k] == '1' {
			cur++
		}

		if cur == 0 || cur == 2 {
			ret = "0" + ret
		} else {
			ret = "1" + ret
		}

		if cur >= 2 {
			cur = 1
		} else {
			cur = 0
		}
	}

	if cur > 0 {
		ret = "1" + ret
	}
	return ret
}
