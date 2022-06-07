package main

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if i == 3 {
			arr = append(arr, 6)
		}
		println(arr[i])
	}
}

func t() {

}
