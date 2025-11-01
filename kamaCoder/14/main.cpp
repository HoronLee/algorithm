#include <bits/stdc++.h>
using namespace std;
int main() {
  int n;
  string s;
  cin >> n;
  getchar(); // 吸收一个回车，因为输入n之后，要输入一个回车
  while (n--) {
    getline(cin, s); // 接收新的一行
    string result;
    // 如果首字母是小写，转换成大写
    if (s[0] >= 'a' && s[0] <= 'z') {
      s[0] -= 32;
    }
    // 将首字母拼接到结果集中
    result += s[0];
    // 遍历整个字符串
    for (int i = 1; i < s.size() - 1; i++) {
      // 如果当前字符是空格，下一个字符不是空格，说明下一个字符是新单词的首字母
      if (s[i] == ' ' && s[i + 1] != ' ') {
        // 判定新单词的首字母是否是小写，小写则转换成大写
        if (s[i + 1] >= 'a' && s[i + 1] <= 'z') {
          s[i + 1] -= 32;
        }
        // 将新单词首字母拼接到结果集中
        result += s[i + 1];
      }
    }
    cout << result << endl;
  }
}
