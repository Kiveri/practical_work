package main

import "fmt"

func search(arr []int, req int) int {
	for ind, elem := range arr {
		if elem == req {

			return ind
		}
	}
	return -1
}

func main() {
	var req int
	arr := []int{13, 22, 31, 14, 65, 456, 37}
	fmt.Println("Искомое значение:")
	fmt.Scan(&req)
	res := search(arr, req)
	fmt.Println("Индекс значения -", res)
}
