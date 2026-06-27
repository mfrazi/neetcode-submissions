func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i:=2; i<len(cost); i++ {
		a, b = b, min(cost[i]+a, cost[i]+b)
	}
	return min(a, b)
}
