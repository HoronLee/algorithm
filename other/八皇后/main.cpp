#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
private:
    // 最终结果集，存储所有有效的棋盘布局
    vector<vector<string>> result;
    
    // 状态数组，用于O(1)时间复杂度的冲突检测
    vector<bool> cols;     // 列是否被占用
    vector<bool> diag1;    // 主对角线 (row - col)
    vector<bool> diag2;    // 副对角线 (row + col)
    int N;

public:
    /**
     * @brief 解决N皇后问题的主函数
     * @param n 棋盘大小
     * @return 包含所有解的二维字符串向量
     */
    vector<vector<string>> solveNQueens(int n) {
        this->N = n;
        // 清空结果集（如果Solution对象被复用）
        result.clear();
        
        // 初始化棋盘，全部填充为 '.'
        vector<string> board(n, string(n, '.'));
        
        // 初始化状态数组
        cols.resize(n, false);
        diag1.resize(2 * n - 1, false);
        diag2.resize(2 * n - 1, false);
        
        // 从第0行开始回溯
        backtrack(0, board);
        
        return result;
    }

private:
    /**
     * @brief 核心回溯函数
     * @param row 当前正在处理的行
     * @param board 当前的棋盘状态
     */
    void backtrack(int row, vector<string>& board) {
        // 基准情况：如果成功处理完所有行（0到N-1）
        // 意味着我们找到了一个完整的解
        if (row == N) {
            result.push_back(board); // 将当前棋盘加入结果集
            return;
        }

        // 遍历当前行的每一列
        for (int col = 0; col < N; ++col) {
            // 计算两个对角线的索引
            // (row - col) 的范围是 [-(N-1), N-1], 加上 (N-1) 偏移量使其变为 [0, 2N-2]
            int diag1_idx = row - col + N - 1;
            // (row + col) 的范围是 [0, 2N-2]
            int diag2_idx = row + col;

            // 剪枝：检查当前位置(row, col)是否安全
            // 如果列、主对角线、副对角线都没有被占用
            if (!cols[col] && !diag1[diag1_idx] && !diag2[diag2_idx]) {
                
                // 1. 做出选择 (Choose)
                board[row][col] = 'Q';
                cols[col] = true;
                diag1[diag1_idx] = true;
                diag2[diag2_idx] = true;

                // 2. 探索下一行 (Explore)
                backtrack(row + 1, board);

                // 3. 撤销选择 (Unchoose) - 关键的回溯步骤
                board[row][col] = '.';
                cols[col] = false;
                diag1[diag1_idx] = false;
                diag2[diag2_idx] = false;
            }
        }
        // 如果for循环结束都没有找到可放的位置，
        // 函数会自动返回到上一层（row - 1），上一层会继续尝试下一列
    }
};

/**
 * @brief 辅助函数，用于打印解决方案
 */
void printSolutions(const vector<vector<string>>& solutions) {
    cout << "找到了 " << solutions.size() << " 种解法。" << endl << endl;
    
    // 只打印前几个解作为示例，否则92个解太多了
    int count = 0;
    for (const auto& board : solutions) {
        count++;
        cout << "--- 解法 " << count << " ---" << endl;
        for (const string& row : board) {
            cout << row << endl;
        }
        cout << endl;
        
        // if (count >= 5) {
        //     cout << "... (及其他解) ..." << endl;
        //     break;
        // }
    }
}

// 主函数
int main() {
    int n = 8; // 解决八皇后问题
    Solution sol;
    vector<vector<string>> solutions = sol.solveNQueens(n);
    
    // 打印结果
    printSolutions(solutions);
    
    return 0;
}
