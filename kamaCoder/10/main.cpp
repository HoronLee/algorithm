// 平均绩点
#include <iostream>
#include <stdio.h>
#include <string>
using namespace std;
int main() {
    string s;
    while (getline(cin, s)) { // 接受一整行字符串
        float sum = 0; // 定义总分数为浮点类型
        int count = 0;
        int flag = 1; // 初始化一个 "真令牌"
      // 遍历字符串，根据不同的分数进行处理
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == 'A') {sum += 4; count++;}
            else if (s[i] == 'B') {sum += 3; count++;}
            else if (s[i] == 'C') {sum += 2; count++;}
            else if (s[i] == 'D') {sum += 1; count++;}
            else if (s[i] == 'F') {sum += 0; count++;}
            else if (s[i] == ' ') continue;
          // 如果不是{A, B, C, D, F}的字符，退出此次循环
            else {
                flag = 0;
                cout << "Unknown" << endl;
                break;
            }
        }
      // 格式化输出
        if (flag) printf("%.2f\n", sum / count);
    }
    return 0;
}
