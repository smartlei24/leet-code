package questions

import "testing"

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit = profit + prices[i] - prices[i-1]
		}
	}

	return profit
}

func TestMaxProfit(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	var profit = maxProfit(prices)
	t.Logf("prices is equal to %d", profit)
}
