package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// bufio中读取数据的接口，因为数据卡的比较严，导致使用fmt.Scan会超时
	scanner := bufio.NewScanner(os.Stdin)
	// 获取数组大小
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 获取数组元素的同时计算前缀和，一般建议切片开大一点防止各种越界问题
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
		if i != 0 {
			arr[i] += arr[i-1]
		}
	}
	/*
	   区间[l, r]的和可以使用区间[0, r]和[0, l - 1]相减得到，
	   在代码中即为arr[r]-arr[l-1]。这里需要注意l-1是否越界
	*/
	for {
		var l, r int
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		if err != nil {
			return
		}
		if l > 0 {
			fmt.Println(arr[r] - arr[l-1])
		} else {
			fmt.Println(arr[r])
		}
	}
}

// func main() {
// 	array, intervals := ReadInput()
// 	// 这里调用你的算法部分
// 	result := quJianHe(array, intervals)
// 	for _, r := range result {
// 		fmt.Println(r)
// 	}
// }

// func quJianHe(array []int, intervals [][2]int) []int {
// 	// TODO: 实现区间和算法
// 	var v, p []int
// 	// 计算前缀
// 	for _, i := range array {
// 		if len(p) == 0 {
// 			p = append(p, i)
// 		} else {
// 			p = append(p, p[len(p)-1]+i)
// 		}
// 	}
// 	// 计算区间和
// 	for _, i := range intervals {
// 		l, r := i[0], i[1]
// 		if l == 0 {
// 			v = append(v, p[r])
// 		} else {
// 			v = append(v, p[r]-p[l-1])
// 		}
// 	}
// 	return v
// }

// // ReadInput 负责处理全部输入，返回整数数组和区间
// func ReadInput() ([]int, [][2]int) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	// 第一行为数组长度
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())
// 	array := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		scanner.Scan()
// 		array[i], _ = strconv.Atoi(scanner.Text())
// 	}
// 	var intervals [][2]int
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		split := strings.Fields(line)
// 		if len(split) == 2 {
// 			l, _ := strconv.Atoi(split[0])
// 			r, _ := strconv.Atoi(split[1])
// 			intervals = append(intervals, [2]int{l, r})
// 		}
// 	}
// 	return array, intervals
// }
