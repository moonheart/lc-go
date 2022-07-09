package main

type numInfo struct {
	count int
	start int
	end   int
}

func findShortestSubArray(nums []int) int {
	m := map[int]*numInfo{}
	maxCount := 0
	maxs := []*numInfo{}
	for i, num := range nums {
		if info, ok := m[num]; !ok {
			info = &numInfo{0, i, i}
			m[num] = info
		}
		m[num].count++
		if m[num].end < i {
			m[num].end = i
		}
		if m[num].count > maxCount {
			maxCount = m[num].count
			maxs = []*numInfo{m[num]}
		} else if m[num].count == maxCount {
			maxs = append(maxs, m[num])
		}
	}
	minLen := len(nums)
	for _, max := range maxs {
		curLen := max.end - max.start + 1
		if curLen < minLen {
			minLen = curLen
		}
	}
	return minLen
}
