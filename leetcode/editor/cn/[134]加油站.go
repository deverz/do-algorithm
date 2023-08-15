package main

// leetcode submit region begin(Prohibit modification and deletion)
func canCompleteCircuit(gas []int, cost []int) int {
	var sum int
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
	}
	// 如果所有站点加的油量和使用的油量相减的总和小于0，则不可能
	if sum < 0 {
		return -1
	}
	var start int // 起始位置
	var tank int  // 当前油箱容量
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		// 如果从start到i，油箱剩余油量小于0，则从start开始不能走完一圈
		if tank < 0 {
			// 油量清空，从该节点的下一个节点开始重新计算
			tank = 0
			start = i + 1
		}
	}
	if start == len(gas) {
		return 0
	}
	return start
}

//leetcode submit region end(Prohibit modification and deletion)
