package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{1, 0, 1, 4, 0, 5, 1, 4})
}

func moveZeroes(nums []int) {
	slow := 0
	for fast := range nums {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
}
