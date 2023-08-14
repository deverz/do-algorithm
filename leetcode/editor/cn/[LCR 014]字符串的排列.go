package main

// leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)    // s2
	windows := make(map[byte]int) // 窗口

	for i := range s1 {
		need[s1[i]] += 1
	}
	left := 0  // 左指针
	right := 0 // 右指针
	valid := 0 // 是否满足s2某个字符的有效计数

	// 开始遍历
	// 核心在于是否包含子串的全部字符，所以窗口的大小是固定的len(s1)
	for right < len(s2) {
		c := s2[right] // 右指针指向的字符
		right += 1     // 右指针右移
		if _, ok := need[c]; ok {
			windows[c] += 1            // 将需要的字符加入窗口
			if windows[c] == need[c] { // 如果窗口中的字符数和need中需要的字符数匹配，则有效计数+1
				valid += 1
			}
		}
		// 如果窗口比s1长度大，则可以挪动左窗口
		for right-left >= len(s1) {
			// 挪动之前先判断是否找到了子串
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left += 1
			// 如果+1之前的左指针字符在窗口内
			if _, ok := windows[d]; ok {
				if need[d] == windows[d] {
					valid -= 1
				}
				windows[d] -= 1
			}
		}

		// 继续遍历，直到右指针到最后
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
