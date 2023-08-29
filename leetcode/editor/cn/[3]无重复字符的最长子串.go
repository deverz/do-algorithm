package main

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	var left = 0
	var right = 0
	var windows = make(map[byte]int, len(s))

	var length = 0

	for right < len(s) {
		c := s[right]
		right++
		windows[c]++
		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
		// 可能不会进入到windows[c] > 1中
		// 所以在外面记录最长子串
		if length < right-left {
			length = right - left
		}
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)
