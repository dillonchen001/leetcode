package main

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var iterms []int

	res = backTrack(candidates, 0, target, iterms, res)
	return res
}

func backTrack(candidates []int, start int, target int, items []int, res [][]int) [][]int {
	if target < 0 {
		return res
	}

	if target == 0 {
		var tmp []int
		for _, val := range items {
			tmp = append(tmp, val)
		}
		res = append(res, tmp)
		return res
	}

	lc := len(candidates)
	for i := start; i < lc; i++ {
		items = append(items, candidates[i])
		res = backTrack(candidates, i, target-candidates[i], items, res)
		items = items[0 : len(items)-1]
	}
	return res
}
