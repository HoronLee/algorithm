#include <bits/stdc++.h>
using namespace std;
void swap(char &a, char &b) { // 交换两个字符串
  char tmp = a;
  a = b;
  b = tmp;
}
int main() {
  int n;
  cin >> n;
  string s;
  while (n--) {
    cin >> s;                                   // 接收字符串s
    for (int i = 0; i < s.size() - 1; i += 2) { // 在s字符串上原地修改
      swap(s[i], s[i + 1]);                     // 调用函数，完成字符串交换
    }
    cout << s << endl;
  }
}
