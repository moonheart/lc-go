package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) (res []string) {
	d := make([]int, len(score))
	copy(d, score)
	sort.Ints(d)
	m := map[int]int{}
	for i, n := range d {
		m[n] = len(score) - i
	}
	for _, n := range score {
		idx := m[n]
		if idx == 1 {
			res = append(res, "Gold Medal")
		} else if idx == 2 {
			res = append(res, "Silver Medal")
		} else if idx == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(idx))
		}
	}
	return
}
