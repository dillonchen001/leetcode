package main

import "fmt"

func main() {
	prices := []int{2, 1, 4, 5, 2, 9, 7}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	lp := len(prices)
	if lp < 2 {
		return 0
	}

	result := 0
	buy := prices[0]
	for i := 1; i < lp; i++ {
		if prices[i] <= buy {
			buy = prices[i]
		} else if prices[i] > buy {
			if i == lp-1 {
				return result + prices[lp-1] - buy
			}
			for j := i + 1; j < lp; j++ {
				if prices[j] < prices[j-1] {
					result += prices[j-1] - buy
					i = j
					buy = prices[i]
					break
				}

				if j == lp-1 {
					return result + prices[lp-1] - buy
				}

			}
		}
	}

	return result
}
