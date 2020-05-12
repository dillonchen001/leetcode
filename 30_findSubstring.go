package main

import "fmt"

func findSubstring(s string, words []string) []int {
	var res []int

	if s == "" || words == nil || len(words) == 0 {
		return res
	}

	wdNun := len(words)
	wdLen := len(words[0])
	wordsMap := make(map[string]int)

	for _, w := range words {
		wordsMap[w]++
	}

	for i := 0; i < len(s)-wdNun*wdLen+1; i++ {
		tmp := make(map[string]int)
		num := 0
		for num < wdNun {
			word := s[i+num*wdLen : i+(num+1)*wdLen]
			_, ok := wordsMap[word]

			if ok {
				tmp[word]++
				if tmp[word] > wordsMap[word] {
					break
				}
			} else {
				break
			}
			num++
		}

		if num == wdNun {
			res = append(res, i)
		}
	}
	return res
}
