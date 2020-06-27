package main

import (
	"container/list"
	"fmt"
)

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

	disMap := make(map[string]int)
	disMap[beginWord] = 0
	queue := list.New()
	queue.PushBack(beginWord)
	relation := make(map[string][]string)

	for queue.Len() > 0 {
		last := queue.Front()
		tmpStr := last.Value.(string)
		queue.Remove(last)
		minDis := disMap[tmpStr]

		neighbors := getNeighbors(tmpStr, wordMap)

		for neighbor, _ := range neighbors {
			isNew := false
			if _, ok := relation[neighbor]; !ok {
				isNew = true
			} else {
				if _, ok := disMap[neighbor]; !ok {
					queue.PushBack(neighbor)
					isNew = true
				} else {
					if minDis+1 < disMap[neighbor] {
						isNew = true
					}
				}
			}

			if isNew {
				tmpSli := []string{tmpStr}
				relation[neighbor] = tmpSli
			} else {
				if minDis+1 == disMap[neighbor] {
					relation[neighbor] = append(relation[neighbor], tmpStr)
				}
			}

			if _, ok := disMap[neighbor]; !ok {
				queue.PushBack(neighbor)
				disMap[neighbor] = minDis + 1
			} else {
				if minDis+1 < disMap[neighbor] {
					disMap[neighbor] = minDis + 1
				}
			}
		}

		delete(disMap, tmpStr)
		delete(wordMap, tmpStr)
	}

	var tmp []string
	tmp = append(tmp, endWord)
	dfs(endWord, tmp, relation, &result, beginWord)
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

func dfs(endWord string, tmp []string, relation map[string][]string, result *[][]string, beginWord string) {
	if endWord == beginWord {
		var sliTmp []string
		for i := len(tmp) - 1; i >= 0; i-- {
			sliTmp = append(sliTmp, tmp[i])
		}
		*result = append(*result, sliTmp)
		return
	}

	if _, ok := relation[endWord]; !ok {
		return
	}

	pre := relation[endWord]
	for _, v := range pre {
		tmp = append(tmp, v)
		dfs(v, tmp, relation, result, beginWord)
		tmp = tmp[0 : len(tmp)-1]
	}
}
