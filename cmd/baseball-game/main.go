package main

import "strconv"

func calPoints(ops []string) (res int) {
	arr := []int{}
	for _, op := range ops {
		switch op {
		case "+":
			arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
			break
		case "D":
			arr = append(arr, arr[len(arr)-1]*2)
			break
		case "C":
			arr = arr[:len(arr)-1]
			break
		default:
			n, _ := strconv.Atoi(op)
			arr = append(arr, n)
		}
	}
	for _, i := range arr {
		res += i
	}
	return
}
