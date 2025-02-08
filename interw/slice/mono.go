package main

import (
	"fmt"
)

func main() {
	s1 := []int64{1, 2, 7, 8, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	s2 := []int64{10, 10, 8, 7, 6, 6, 4, 3, 1}
	s3 := []int64{1, 2, 2, 5, 5, 1, 1, 2}
	s4 := []int64{1}

	fmt.Println("s1:", mono(s1))
	fmt.Println("s2:", mono(s2))
	fmt.Println("s3:", mono(s3))
	fmt.Println("s4:", mono(s4))
}

func mono(s []int64) bool {
	ch1 := 0
	ch2 := 0
	for i := 0; i < len(s)-1; i++ {
		curr := s[i]
		next := s[i+1]
		if curr <= next {
			ch1++
		}
		if curr >= next {
			ch2++
		}
	}
	if ch1 == len(s)-1 || ch2 == len(s)-1 {
		return true
	}

	return false
}
