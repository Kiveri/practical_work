package main

import "fmt"

func dubFind(arr []int) []int {
	var res []int
	for _, elem := range arr {
		notFound := false
		for _, elem2 := range res {
			if elem2 == elem {
				notFound = true
				break
			}
		}
		if !notFound {
			res = append(res, elem)
		}
	}

	return res
}

func main() {
	arr := []int{1, 1, 2, 3, 3, 3, 4}

	fmt.Println(arr)
	fmt.Println(dubFind(arr))
}
