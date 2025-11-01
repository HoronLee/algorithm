// 有序数组的平方
package main

import (
	"fmt"
	"sort"
)

func main() {
	ret := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Printf("%v", ret)
}

// 双指针方法
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i, j := 0, len(nums)-1
	k := len(result) - 1
	for k >= 0 {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			result[k] = nums[j] * nums[j]
			j--
		} else {
			result[k] = nums[i] * nums[i]
			i++
		}
		k--
	}
	return result
}

func carl_sortedSquares(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			ans[k] = lm
			i++
		} else {
			ans[k] = rm
			j--
		}
		k--
	}
	return ans
}

// 暴力方法
func SortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}
