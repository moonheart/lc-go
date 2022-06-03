package main

func firstUniqChar(s string) byte {
	// 利用题目条件 仅包含小写字母, 初始化一个长度 26 的数组
	// 数组的 index 用于区分字母, 数组的 value 用于记录字母首次出现的位置
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		// 由于我们是要寻找第一个出现的字符, index 要尽量小, 所以设置 len(s) 作为默认 index
		pos[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n {
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	// 寻找最小的 index
	ans := n
	for _, p := range pos[:] {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return s[ans]
	}
	return ' '
}

func main() {
	firstUniqChar("leetcodelike")
}
