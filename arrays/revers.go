package main

import "fmt"

func rvrsArr(arr []int) []int {
	n := len(arr)
	rvrs := make([]int, n)
	for i := 0; i < n; i++ {
		rvrs[i] = arr[n-i-1]
	}
	return rvrs
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resArr := rvrsArr(arr)
	fmt.Println(resArr)
}
