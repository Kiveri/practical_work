package main

import (
	"fmt"
)

func main() {
	a() // 0123 01234
}

func a() {
	x := []int{}
	x = append(x, 0)  // 1-1	0
	x = append(x, 1)  // 2-2	01
	x = append(x, 2)  // 3-4	012
	y := append(x, 3) // 4-4 0123
	z := append(y, 4) // 5-8 01234
	fmt.Println(y, z)
}
