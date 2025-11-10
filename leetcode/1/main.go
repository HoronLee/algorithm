package main

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if index, isHave := m[diff]; isHave {
			return []int{index, i}
		}
		m[num] = i
	}
	return nil
}
