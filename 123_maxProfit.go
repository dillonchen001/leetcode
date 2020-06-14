package main

import (
	"fmt"
)

func main() {
	prices := []int{2, 1, 4, 5, 2, 9, 7}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	lp := len(prices)
	if lp < 2 {
		return 0
	}

	var profits [2]int
	buy := prices[0]
	for i := 1; i < lp; i++ {
		if prices[i] <= buy {
			buy = prices[i]
		} else if prices[i] > buy {
			if i == lp-1 {
				profits = getAllProfits(profits, prices[lp-1]-buy)
				break
			}
			for j := i + 1; j < lp; j++ {
				if prices[j] < prices[j-1] {
					profits = getAllProfits(profits, prices[j-1]-buy)
					i = j
					buy = prices[i]
					break
				}

				if j == lp-1 {
					profits = getAllProfits(profits, prices[lp-1]-buy)
					i = lp
				}

			}
		}
	}

	return profits[0] + profits[1]
}

func getAllProfits(profits [2]int, profit int) [2]int {
	//为空
	if profits[0] == 0 {
		profits[0] = profit
		return profits
	}

	if profits[1] == 0 {
		profits[1] = profit
		return profits
	}

	if profits[1] < profits[0] {
		profits[0], profits[1] = profits[1], profits[0]
	}

	if profit <= profits[0] {
		return profits
	}
	profits[0] = profit
	return profits
}
