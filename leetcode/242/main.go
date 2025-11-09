package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var records = [26]int{}
	for index := 0; index < len(s); index++ {
		if s[index] == t[index] {
			continue
		}
		sCharIndex := s[index] - 'a'
		records[sCharIndex]++
		tCharIndex := t[index] - 'a'
		records[tCharIndex]--
	}
	for _, record := range records {
		if record != 0 {
			return false
		}
	}
	return true
}

// func isAnagram(s string, t string) bool {
// 	record := [26]int{}

// 	for _, r := range s {
// 		record[r-rune('a')]++
// 	}
// 	for _, r := range t {
// 		record[r-rune('a')]--
// 	}

// 	return record == [26]int{} // record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
// }
