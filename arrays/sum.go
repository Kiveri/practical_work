package main

import "fmt"

func sumNum(numbers []int) int {
	sum := 0
	for _, elem := range numbers {
		sum += elem
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := sumNum(numbers)
	fmt.Println(res)
}
