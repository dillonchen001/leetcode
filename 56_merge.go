package main

type ByStart [][]int

func (intervals ByStart) Len() int {
	return len(intervals)
}

func (intervals ByStart) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func (intervals ByStart) Less(i, j int) bool {
	return intervals[i][0] < intervals[j][0]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(ByStart(intervals))

	li := len(intervals)
	if li <= 1 {
		return intervals
	}

	var ret [][]int
	var tmp int
	var j int
	for i := 0; i < li; i++ {
		tmp = intervals[i][1]
		j = i + 1
		for j < li && intervals[j][0] <= tmp {
			if intervals[j][1] > tmp {
				tmp = intervals[j][0]
			}
			j++
		}
		ret = append(ret, []int{intervals[i][0], tmp})
		i = j
	}
	return ret
}
