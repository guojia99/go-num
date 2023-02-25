package _22

// 给你一个整数数组 prices ，其中prices[i] 表示某支股票第 i 天的价格。
//
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润。
//
// 示例 1：
//
// 输入：prices = [7,1,5,3,6,4]
// 输出：7
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
//
//	随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
//	 总利润为 4 + 3 = 7 。
//
// 示例 2：
//
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
//
//	总利润为 4 。
//
// 示例3：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。

func maxProfitDp(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// sell：持有现金
	// buy：持有股票
	// 状态转移：buy → sell → buy → sell → buy → sell → ...
	var (
		sell = make([]int, len(prices))
		buy  = make([]int, len(prices))
	)
	sell[0], buy[0] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]) // 总和上一次的非buy/sell的状态做对比
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
	}
	return sell[len(prices)-1]
}

// 贪心
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	var ans int

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
