package main

import (
	"fmt"
	"strconv"
)

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	lp := len(prices)

	if lp < 2 {
		return 0
	}

	maxK := 2

	if lp/2 < maxK {
		maxK = lp / 2
	}

	dp := make(map[string]int)

	for i := 0; i < lp; i++ {
		for k := maxK; k > 0; k-- {
			if i == 0 {
				dp["0-"+strconv.Itoa(k)+"-0"] = 0
				dp["0-"+strconv.Itoa(k)+"-1"] = 0 - prices[0]
			} else {
				dp[strconv.Itoa(i)+"-"+strconv.Itoa(k)+"-0"] = getMax(dp[strconv.Itoa(i-1)+"-"+strconv.Itoa(k)+"-0"], dp[strconv.Itoa(i-1)+"-"+strconv.Itoa(k)+"-1"]+prices[i])
				dp[strconv.Itoa(i)+"-"+strconv.Itoa(k)+"-1"] = getMax(dp[strconv.Itoa(i-1)+"-"+strconv.Itoa(k)+"-1"], dp[strconv.Itoa(i-1)+"-"+strconv.Itoa(k-1)+"-0"]-prices[i])
			}
		}
	}

	fmt.Println(dp)
	return dp[strconv.Itoa(lp-1)+"-"+strconv.Itoa(maxK)+"-0"]
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
