package main

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	var items []int
	result = backtrace(candidates, 0, target, items, result)
	return result
}

func backtrace(candidates []int, start int, target int, items []int, res [][]int) [][]int {
	if target < 0 {
		return res
	}

	if target == 0 {
		var tmp []int
		for _, value := range items {
			tmp = append(tmp, value)
		}
		res = append(res, tmp)
		return res
	}

	lc := len(candidates)
	for i := start; i < lc; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		items = append(items, candidates[i])
		res = backtrace(candidates, i+1, target-candidates[i], items, res)
		items = items[0 : len(items)-1]
	}
	return res
}
