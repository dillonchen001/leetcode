package main

func maxProfit(prices []int) int {
	ln := len(prices)
	if ln < 2 {
		return 0
	}
	min := prices[0]
	result := 0
	for i := 0; i < ln; i++ {
		if prices[i]-min > result {
			result = prices[i] - min
		}

		if min > prices[i] {
			min = prices[i]
		}
	}
	return result
}
