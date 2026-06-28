func coinChange(coins []int, amount int) int {
	INF := amount+1
	dp := make([]int, amount+1)
	for i:=1; i<=amount; i++ {
		dp[i] = INF
	}

	for i:=1; i<=amount; i++ {
		for j:=0; j<len(coins); j++ {
			if i >= coins[j] &&
			 dp[i] > dp[i - coins[j]] + 1 {
				dp[i] = dp[i - coins[j]] + 1
			}
		}
	}
	
	if dp[amount] == INF {
		return -1
	}

	return dp[amount]
}

//      1 5 6 10
// 1    1 0 0  0
// 2    2 0 0  0
// 3    3 0 0  0
// 4    4 0 0  0
// 5    0 1 0  0
// 6    0 0 1  0
// 7    1 0 1  0
// 8    2 0 1  0
// 9    3 0 1  0
// 10   0 0 0  1
// 11   1 0 0  1
// 12   0 0 2  0


//      1  5 6 10
// 1    1  0 0  0
// 2    2  0 0  0
// 3    3  0 0  0
// 4    4  0 0  0
// 5    0  1 0  0
// 6    6  0 1  0
// 7    7  0 0  0
// 8    8  0 0  0
// 9    9  0 0  0
// 10   10 2 0  1
// 11   11 0 0  0
// 12   12 0 2  0