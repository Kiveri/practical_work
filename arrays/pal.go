package main

import "fmt"

func isPal(arr []int) bool {
	l := 0
	r := len(arr) - 1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 2, 1}
	fmt.Println(isPal(arr))
}
