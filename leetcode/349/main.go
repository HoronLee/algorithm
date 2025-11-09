package main

func intersection(nums1 []int, nums2 []int) []int {
	// 使用空值来模拟set（不重复集合）
	set := make(map[int]struct{})
	res := make([]int, 0)
	// 如果set中不存在nums1中的值则添加一个记录
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		// 如果在set中发现了nums2中的键，则添加到结果数组中，并且要删掉这个键
		if _, ok := set[v]; ok {
			res = append(res, v)
			// 防止冲突添加到结果数组（题目要求）
			delete(set, v)
		}
	}
	return res
}
