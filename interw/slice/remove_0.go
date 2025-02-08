package main

import (
	"fmt"
)

func main() {
	s1 := []int64{1, 2, 0, 0, 4, 5, 0}
	s2 := []int64{0}
	var s0 []int64

	s1r := remove(s1)
	fmt.Println(s1r)

	s2r := remove(s2)
	fmt.Println(s2r)

	s0r := remove(s0)
	fmt.Println(s0r)
}

func remove(s []int64) []int64 {
	res := make([]int64, 0)
	for _, v := range s {
		if v != 0 {
			res = append(res, v)
		}
	}

	return res
}
