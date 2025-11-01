package main

import "fmt"

func main() {
	// 设置 N=8 (八皇后问题)
	const N = 8

	// 初始化棋盘：
	// 创建一个 N x N 的二维切片，默认所有值为 0
	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}

	// 从第 0 行开始解决问题
	solveNQueens(board, 0, N)
}

// isSafe 检查在 (row, col) 位置放皇后是否安全
// N 是棋盘的大小 (N=8)
func isSafe(board [][]int, row int, col int, N int) bool {
	// 1. 检查这一列的上方 (⬆️)
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	// 2. 检查左上对角线 (↖️)
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// 3. 检查右上对角线 (↗️)
	for i, j := row-1, col+1; i >= 0 && j < N; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// 如果所有检查都通过了
	return true
}

// printSolution 打印棋盘
func printSolution(board [][]int, N int) {
	fmt.Println("找到一个解：")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 1 {
				fmt.Print("Q ") // 皇后
			} else {
				fmt.Print(". ") // 空格
			}
		}
		fmt.Println() // 换行
	}
	fmt.Println("--------------------")
}

// solveNQueens 递归回溯函数
// (我们一步步构建的函数)
func solveNQueens(board [][]int, row int, N int) {
	// 1. 退出条件 (Base Case): 成功处理完所有行
	if row == N {
		printSolution(board, N)
		// 注意：这里我们只返回，以便回溯可以继续寻找所有解
		// 如果只想找一个解，可以在这里用 os.Exit(0)
		return
	}

	// 2. 递归步骤: 遍历当前行的所有列
	for col := 0; col < N; col++ {
		// 检查 (row, col) 是否安全
		if isSafe(board, row, col, N) {

			// (1) 做出选择：放置皇后
			board[row][col] = 1

			// (2) 探索：递归到下一行
			solveNQueens(board, row+1, N)

			// (3) 撤销选择 (Backtrack)：
			// 把皇后拿走，尝试这一行的下一列
			board[row][col] = 0
		}
	}
}
