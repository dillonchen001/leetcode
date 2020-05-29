package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) < 2 {
		return intervals
	}

	var ret [][]int
	li := len(intervals)

	if li < 1 {
		ret = append(ret, newInterval)
		return ret
	}

	i := 0
	for i < li && intervals[i][1] < newInterval[0] {
		ret = append(ret, intervals[i])
		i++
	}

	for i < li && intervals[i][0] <= newInterval[1] {
		if newInterval[0] > intervals[i][0] {
			newInterval[0] = intervals[i][0]
		}

		if newInterval[1] < intervals[i][1] {
			newInterval[1] = intervals[i][1]
		}
		i += 1
	}
	ret = append(ret, newInterval)

	for i < li {
		ret = append(ret, intervals[i])
		i++
	}
	return ret
}
