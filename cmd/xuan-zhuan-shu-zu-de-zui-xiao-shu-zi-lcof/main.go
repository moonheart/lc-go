package main

func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		middle := (high-low)/2 + low
		if numbers[middle] < numbers[high] {
			high = middle
		} else if numbers[middle] > numbers[high] {
			low = middle + 1
		} else {
			high--
		}
	}
	return numbers[low]
}
