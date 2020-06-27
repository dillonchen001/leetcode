package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var result [][]string
	wordMap := make(map[string]bool)
	lw := len(beginWord)
	exist := false
	for _, v := range wordList {
		if v == endWord {
			exist = true
		}

		if len(v) != lw {
			continue
		}
		wordMap[v] = true
	}

	if !exist {
		return result
	}

	tmp := make(map[string]int)
	tmp[beginWord] = 0
	bfs(beginWord, endWord, wordMap, tmp, &result)
	return result
}

func getNeighbors(word string, wordMap map[string]bool) map[string]bool {
	result := make(map[string]bool)
	lw := len(word)
	for ch := 'a'; ch <= 'z'; ch++ {
		for i := 0; i < lw; i++ {
			if word[i] == byte(ch) {
				continue
			}
			tmp := word[0:i] + string(byte(ch)) + word[i+1:lw]

			if _, ok := wordMap[tmp]; ok {
				result[tmp] = true
			}
		}
	}
	return result
}

func bfs(beginWord string, endWord string, wordMap map[string]bool, tmp map[string]int, result *[][]string) {
	neighbors := getNeighbors(beginWord, wordMap)

	if len(neighbors) == 0 {
		return
	}

	if _, ok := neighbors[endWord]; ok {
		seri := len(tmp)
		tmp[endWord] = seri

		if *result != nil {
			if len((*result)[0]) > len(tmp) {
				tmp := (*result)[0:0]
				result = &tmp
			} else if len((*result)[0]) < len(tmp) {
				delete(tmp, endWord)
				return
			}
		}

		slTmp := make([]string, len(tmp))
		for key, val := range tmp {
			slTmp[val] = key
		}

		*result = append(*result, slTmp)
		delete(tmp, endWord)
		return
	}

	for key, _ := range neighbors {
		if _, ok := tmp[key]; ok {
			continue
		}

		seri := len(tmp)
		tmp[key] = seri
		bfs(key, endWord, wordMap, tmp, result)
		delete(tmp, key)
	}
}
