package main

import "fmt"

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 2; j >= i; j-- {
				arr[j+1] = arr[j]
			}
			i++
		}
	}
}

func duplicateZeros2(arr []int) {
	cnt0 := 0
	for _, i := range arr {
		if i == 0 {
			cnt0++
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			cnt0--
		}
		if i+cnt0 < len(arr) {
			arr[i+cnt0] = arr[i]
			if arr[i] == 0 && i+cnt0+1 < len(arr) {
				arr[i+cnt0+1] = 0
			}
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Println(arr)
	duplicateZeros(arr)
	fmt.Println(arr)
}
