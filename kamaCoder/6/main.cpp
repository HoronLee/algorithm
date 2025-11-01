// 数组的倒序与隔位输出
#include <iostream>
#include <vector>
using namespace std;
int main() {
  int n, num;
  cin >> n;
  vector<int> nums;
  while (n--) {
    cin >> num;
    nums.push_back(num);
  }
  for (int i = nums.size() - 1; i >= 0; i--) {
    cout << nums[i];
    if (i > 0) {
      cout << " ";
    }
  }
  cout << endl;
  for (int i = 0; i < nums.size(); i += 2) {
    cout << nums[i];
    if (i < nums.size() - 1) {
      cout << " ";
    }
  }
  return 0;
}
