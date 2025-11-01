// 开发商购买土地
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int // 矩阵的高和宽
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	params := strings.Split(line, " ")
	n, _ = strconv.Atoi(params[0])
	m, _ = strconv.Atoi(params[1])
	land := make([][]int, n)
	// 构建矩阵
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")
		land[i] = make([]int, m)
		for j := 0; j < m; j++ {
			value, _ := strconv.Atoi(values[j])
			land[i][j] = value
		}
	}
	preMatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		preMatrix[i] = make([]int, m+1)
	}
	// 计算前缀和矩阵
	for a := 1; a < n+1; a++ {
		for b := 1; b < m+1; b++ {
			preMatrix[a][b] = land[a-1][b-1] + preMatrix[a-1][b] + preMatrix[a][b-1] - preMatrix[a-1][b-1]
		}
	}
	totalSum := preMatrix[n][m]
	minDiff := math.MaxInt32
	// 行
	for i := 1; i < n; i++ {
		topSum := preMatrix[i][m]
		bottomSum := totalSum - topSum
		diff := int(math.Abs(float64(topSum - bottomSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}
	// 列
	for j := 1; j < m; j++ {
		topSum := preMatrix[n][j]
		bottomSum := totalSum - topSum
		diff := int(math.Abs(float64(topSum - bottomSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}
	fmt.Println(minDiff)
}
