package main

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
OUT:
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				continue OUT
			}
		}
		res++
	}
	return res
}

//
//func findTheDistanceValue1(arr1 []int, arr2 []int, d int) int {
//	res := 0
//	sort.Ints(arr1)
//	sort.Ints(arr2)
//	startJ := sort.SearchInts(arr2, arr1[0])
//	l := len(arr2)
//	for i, i2 := range arr1 {
//		f := true
//		if startJ < l {
//			f = f && (abs(i2-arr2[startJ]) > d)
//		}
//		if startJ >=
//		if i2-arr2[i] > d {
//			res++
//		}
//	}
//	return res
//}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
