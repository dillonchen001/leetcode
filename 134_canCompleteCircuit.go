package main

func canCompleteCircuit(gas []int, cost []int) int {
	lg := len(gas)

	start := 0
	totalTank := 0
	currTank := 0
	for i := 0; i < lg; i++ {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]

		if currTank < 0 {
			start = i + 1
			currTank = 0
		}
	}

	if totalTank >= 0 {
		return start
	}
	return -1
}
