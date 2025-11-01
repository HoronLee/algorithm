// 螺旋矩阵II
package main

import (
	"fmt"
)

func main() {
	ret := generateMatrix(4)
	fmt.Printf("%v", ret)
}

// 思路：可以设置一个offset变量，表示当前层距离边界的偏移量
func generateMatrix(n int) [][]int {
	// 循环采用左闭右开区间
	res := [][]int{}
	for range n {
		res = append(res, make([]int, n))
	}
	startx, starty := 0, 0
	loop, mid := n/2, n/2
	count, offset := 1, 1
	var i, j int
	for loop > 0 {
		i = startx
		j = starty
		// 四个循环模拟了螺旋矩阵的四个边
		// ➡️
		for ; j < n-offset; j++ {
			res[i][j] = count
			count += 1
		}
		// ⬇️
		for ; i < n-offset; i++ {
			res[i][j] = count
			count += 1
		}
		// ⬅️
		for ; j > starty; j-- {
			res[i][j] = count
			count += 1
		}
		// ⬆️
		for ; i > startx; i-- {
			res[i][j] = count
			count += 1
		}
		startx += 1
		starty += 1
		offset += 1
		loop--
	}
	if n%2 == 1 {
		res[mid][mid] = count
	}
	return res
}
