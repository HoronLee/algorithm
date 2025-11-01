// 二分查找
package main

import (
	"fmt"
)

func main() {
	ret := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Printf("%v", ret)
}

// 左右都为闭区间
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1 // 等于整除2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
