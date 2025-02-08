package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(uniqRand(10))
}

func uniqRand(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(10)
	}

	return s
}
