package main

import "fmt"

func main() {
	ret := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Printf("%v", ret)
}

// 双指针，滑动窗口
func minSubArrayLen(target int, nums []int) int {
	// 窗口式满足条件的子数组
	// 解题的关键在于 窗口的起始位置如何移动
	result := len(nums) + 1
	// i 为滑动窗口的前端
	var sum, i, subLen int
	for j := range nums {
		sum += nums[j]
		for sum >= target {
			subLen = j - i + 1
			if result > subLen {
				result = subLen
			}
			// 关键步骤，保持滑动窗口的右端 j 不变
			// 将窗口左侧向右边缩一次，也就是减掉索引 i 的值
			// 然后再次对比是否为满足条件的子串
			sum -= nums[i]
			i++
		}
	}
	// 找不到符合条件的子串需要返回 0
	if result == len(nums)+1 {
		return 0
	}
	return result
}

func MinSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
