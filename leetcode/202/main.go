package main

func isHappy(n int) bool {
	// 模拟set
	set := make(map[int]bool)
	// 如果不为1并且set中不存在曾经计算出的sum值（即保证不陷入死循环）
	for n != 1 && !set[n] {
		// 获取新的sum
		// 并且标记上一个sum为出现过的
		n, set[n] = getSum(n), true
	}
	// 如果是1则为快乐数
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
