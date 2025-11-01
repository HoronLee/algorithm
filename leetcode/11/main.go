/*
 * 贪心算法，只让较短板向中心寻找可能的长板
 */
package main

import "fmt"

func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var res int
	for i < j {
		if height[i] < height[j] {
			res = max(res, height[i]*(j-i))
			i++
		} else {
			res = max(res, height[j]*(j-i))
			j--
		}
	}
	return res
}
