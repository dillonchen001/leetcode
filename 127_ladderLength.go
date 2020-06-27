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

func ladderLength(beginWord string, endWord string, wordList []string) int {
	result := 0
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

		if tmpStr == endWord {
			result = disMap[tmpStr] + 1
		}

		delete(disMap, tmpStr)
		delete(wordMap, tmpStr)
	}
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
