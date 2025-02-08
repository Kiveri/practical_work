package main

import (
	"fmt"
)

func main() {
	s1 := []int64{1, 2, 3, 4}
	s2 := []int64{4, 5, 6, 7, 8}

	fmt.Println(zip(s1, s2))
}

func zip(s1, s2 []int64) [][]int64 {
	res := make([][]int64, 0)
	for i := 0; i < len(s1); i++ {
		res = append(res, []int64{s1[i], s2[i]})
	}

	return res
}
