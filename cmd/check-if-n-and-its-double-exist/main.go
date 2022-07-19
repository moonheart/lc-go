package main

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]*2 == arr[j] || arr[j]*2 == arr[i] {
				return true
			}
		}
	}
	return false
}

func checkIfExist2(arr []int) bool {
	m := map[int]bool{}
	for _, i := range arr {
		if m[i] {
			return true
		}
		m[i*2] = true
		if i&1 == 0 {
			m[i/2] = true
		}
	}
	return false
}
