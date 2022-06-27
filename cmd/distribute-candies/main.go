package main

func distributeCandies(candyType []int) (res int) {
	m := map[int]bool{}
	half := len(candyType) / 2
	for i, n := 0, 0; i < len(candyType) && n < half && res < half; i++ {
		if !m[candyType[i]] {
			m[candyType[i]] = true
			res++
			n++
		}
	}
	return
}

func distributeCandies2(candyType []int) int {
	set := map[int]struct{}{}
	for _, t := range candyType {
		set[t] = struct{}{}
	}
	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}
	return ans
}
