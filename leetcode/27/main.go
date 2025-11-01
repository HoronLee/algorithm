// 移除元素
package main

import (
	"fmt"
)

func main() {
	ret := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Printf("%v", ret)
}

// 快慢指针
func removeElement(nums []int, val int) int {
	slow := 0                // 慢指针，用于构建无目标值的新数组
	for fast := range nums { // 快指针遍历数组
		// 当 fast 指针对应值不是目标值时，复制到 slow 指针处
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow // slow 即为新数组的长度
}

// 相向双指针法
func removeElement1(nums []int, val int) int {
	// 有点像二分查找的左闭右闭区间 所以下面是<=
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		// 各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}
