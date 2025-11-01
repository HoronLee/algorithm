// A+B Question IV
#include <iostream>
using namespace std;
int main() {
  int n;
  while (cin >> n) {
    if (n == 0)
      break;
    int sum = 0;
    int a;
    while (n--) {
      cin >> a;
      sum += a;
    }
    cout << sum << endl;
  }
  return 0;
}
