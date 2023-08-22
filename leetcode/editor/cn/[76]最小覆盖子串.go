package main

import (
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	//[left,right) 左闭右开
	left := 0                     // 左指针
	right := 0                    // 右指针
	windows := make(map[byte]int) // 窗口
	needs := make(map[byte]int)   // 小串需要的字符数统计
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	valid := 0 // 标记是否已经满足小串的字符数需要
	start := 0
	length := math.MaxInt32
	for right < len(s) {
		c := s[right]
		// 右移，右移前先记录需要移动的字符
		right++
		// 该字符是否是子串需要的
		if _, ok := needs[c]; ok {
			// 每移动一个字符，窗口++
			windows[c]++
			if windows[c] == needs[c] {
				valid++
			}
		}

		// 如果已经找全了子串，则开始挪动左指针
		for valid == len(needs) {
			// 如果本次比上次找到的子串长度小，则记录
			if length > right-left {
				start = left
				length = right - left // [left,right)
			}

			// 左指针移动，移动前先记录要移动的字符
			d := s[left]
			left++
			if _, ok := needs[d]; ok {
				if windows[d] == needs[d] {
					valid--
				}
				// 窗口--
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
