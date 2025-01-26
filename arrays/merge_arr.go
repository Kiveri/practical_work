package main

import "fmt"

func mergeArr(arr1 []int, arr2 []int) []int {
	arrmrg := append(arr1, arr2...)
	return arrmrg
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{50, 600, 70}
	res := mergeArr(arr1, arr2)
	fmt.Println(res)
}
