package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0
	endSec := 0
	for _, n := range timeSeries {
		if n > endSec {
			total += duration
		} else {
			total += n + duration - endSec
		}
		endSec = n + duration
	}
	return total
}
