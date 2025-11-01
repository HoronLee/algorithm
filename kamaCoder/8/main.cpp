// 奇怪的信
#include <iostream>
#include <vector>
using namespace std;
int main() {
  int n;
  while (cin >> n) {
    int result = 0;
    while (n != 0) {
      int a = n % 10;
      n /= 10;
      if (a % 2 == 0) {
        result += a;
      }
    }
    cout << result << endl;
    cout << endl;
  }
}
