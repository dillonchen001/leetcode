package main

import (
	"fmt"
)

func main() {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	fmt.Println(fullJustify(words, maxWidth))
}

func fullJustify(words []string, maxWidth int) []string {
	lw := len(words)

	if lw < 1 {
		return words
	}

	var ret []string

	length := 0
	num := 0

	blank := 0
	reminder := 0
	quotient := 0
	last := 0

	for i := 0; i < lw; i++ {
		last = i - 1
		length += len(words[i])
		if length+num < maxWidth {
			num++
			continue
		}

		if length+num == maxWidth {
			num++
			if i == lw-1 {
				break
			}
			last = i
		} else {
			length -= len(words[i])
		}

		tmp := ""
		blank = maxWidth - length

		if num == 1 {
			reminder = 0
			quotient = blank
		} else {
			quotient = blank / (num - 1)
			reminder = blank % (num - 1)
		}

		for j := 0; j < reminder; j++ {
			tmp += words[last-num+j+1]
			for k := 0; k <= quotient; k++ {
				tmp += " "
			}
		}

		m := num - 1
		if num == 1 {
			m = num
		}
		for j := reminder; j < m; j++ {
			tmp += words[last-num+j+1]
			for k := 0; k < quotient; k++ {
				tmp += " "
			}
		}
		if num > 1 {
			tmp += words[last]
		}
		ret = append(ret, tmp)

		if last == i-1 {
			length = len(words[i])
			num = 1
		} else {
			length = 0
			num = 0
		}
	}

	tmp := ""

	for j := num; j > 1; j-- {
		tmp += words[lw-j]
		tmp += " "
	}
	tmp += words[lw-1]

	for k := 0; k < maxWidth-length-num+1; k++ {
		tmp += " "
	}
	ret = append(ret, tmp)
	return ret
}
