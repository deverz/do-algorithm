package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	left := 0
	right := 0
	windows := make(map[byte]int, len(t))
	needs := make(map[byte]int, len(t))
	// 需要匹配的子串的字符数统计
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	valid := 0              // 符合子串的字符要求统计
	start := 0              // 最后匹配到的符合最小子串的左指针索引下表
	length := math.MaxInt32 // 最后匹配到的符合最小子串的长度 right-left

	for right < len(s) {
		// 可能要入滑动窗口的字符
		c := s[right]
		// 右移
		right++
		if _, ok := needs[c]; ok {
			// 符合要求，放入窗口
			windows[c]++
			if needs[c] == windows[c] {
				valid++
			}
		}
		// 左移
		for valid == len(needs) {
			// 重置当前最少匹配长度
			if right-left < length {
				length = right - left
				start = left
			}
			// 可能要出滑动窗口的字符
			d := s[left]
			// 移动
			left++
			if _, ok := needs[d]; ok {
				if needs[d] == windows[d] {
					valid--
				}
				windows[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]

}

//leetcode submit region end(Prohibit modification and deletion)
