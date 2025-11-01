package main

import (
	"fmt"
	"slices"
)

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	// 1. 先对数组排序
	slices.Sort(nums)
	n := len(nums)
	// 2. 遍历数组，固定第一个数 nums[i]
	for i := range nums {
		// 如果 nums[i] > 0，因为数组已排序，后面的数都大于0，不可能相加为0
		if nums[i] > 0 {
			break
		}
		// 3. 对 nums[i] 进行去重，防止出现重复的三元组
		// 确保 i 是第一个不重复的元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 4. 设置左右指针
		L := i + 1
		R := n - 1
		// 5. 关键的内层循环，移动双指针
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			switch {
			case sum == 0:
				// 找到一个解
				res = append(res, []int{nums[i], nums[L], nums[R]})
				// 对 L 和 R 进行去重
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				// 移动指针，寻找下一个可能的解
				L++
				R--
			case sum < 0:
				// 和太小，移动左指针
				L++
			case sum > 0:
				// 和太大，移动右指针
				R--
			}
		}
	}
	return res
}
