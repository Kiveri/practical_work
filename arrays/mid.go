package main

import "fmt"

func midNum(numbers []float32) float32 {
	var sum float32
	for _, elem := range numbers {
		sum += elem
	}
	mid := sum / float32(len(numbers))
	return mid
}

func main() {
	numbers := []float32{1, 2, 3, 4, 5, 6, 7, 8, 11}
	res := midNum(numbers)
	fmt.Println(res)
}
