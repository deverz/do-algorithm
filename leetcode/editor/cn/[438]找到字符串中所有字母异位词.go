package main

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	left := 0
	right := 0
	windows := make(map[byte]int, len(p))
	needs := make(map[byte]int, len(p))
	for i := 0; i < len(p); i++ {
		needs[p[i]]++
	}
	valid := 0

	start := 0
	r := make([]int, 0) // 结果集

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c]++
			if needs[c] == windows[c] {
				valid++
			}
		}
		// 需要匹配到子串全部
		for right-left >= len(p) && valid == len(needs) {
			start = left
			d := s[left]
			left++
			if _, ok := needs[d]; ok {
				if needs[d] == windows[d] {
					// 需要确保，长度一样才属于异位
					if right-start == len(p) {
						r = append(r, start)
					}
					valid--
				}
				windows[d]--
			}
		}
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
