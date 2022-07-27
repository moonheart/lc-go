package main

func isHappy(n int) bool {
	m := map[int]bool{}
	var check func(int) bool
	check = func(i int) bool {
		if m[i] {
			return false
		}
		m[i] = true
		sum := 0
		for i > 0 {
			sum += (i % 10) * (i % 10)
			i = i / 10
		}
		if sum == 1 {
			return true
		}
		return check(sum)
	}
	return check(n)
}

func isHappy1(n int) bool {
	step := func(i int) int {
		sum := 0
		for i > 0 {
			sum += (i % 10) * (i % 10)
			i = i / 10
		}
		return sum
	}
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}
