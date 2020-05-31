package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 1
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 1
	}

	var res []int
	init := 1
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] > 0 {
			init = 0
		}
		res = append(res, init)
	}

	if obstacleGrid[0][0] > 0 {
		return 0
	}

	for j := 1; j < m; j++ {
		for k := 0; k < n; k++ {
			if obstacleGrid[j][k] > 0 {
				res[k] = 0
			} else if k == 0 {
				continue
			} else {
				res[k] += res[k-1]
			}
		}
	}

	return res[n-1]

}
