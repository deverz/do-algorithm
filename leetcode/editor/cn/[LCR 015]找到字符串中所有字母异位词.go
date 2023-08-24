package main

import (
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	// 左指针
	var left = 0
	// 右指针
	var right = 0
	// 窗口
	var windows = make(map[byte]int, len(p))
	// 子串所需字符统计
	var needs = make(map[byte]int, len(p))
	for i := 0; i < len(p); i++ {
		needs[p[i]]++
	}
	// 符合needs要求标记
	var flag = 0
	// 结果集
	var ret = make([]int, 0)

	// 开始挪动右指针
	for right < len(s) {
		// 记录要进窗口的字符
		var in = s[right]
		// 右移
		right++
		if _, ok := needs[in]; ok {
			windows[in]++
			if needs[in] == windows[in] {
				flag++
			}
		}
		// 如果符合needs,开始挪动左指针右移
		for flag == len(needs) && left < right {
			if right-left == len(p) {
				ret = append(ret, left)
			}
			// 记录要出窗口的字符
			out := s[left]

			// 左指针移动
			left++
			if _, ok := needs[out]; ok {
				if needs[out] == windows[out] {
					flag--
				}
				windows[out]--
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
