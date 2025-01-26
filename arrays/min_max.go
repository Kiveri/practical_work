package main

import "fmt"

func minMaxNum(numbers []int) (int, int) {
	var min int
	var max int
	for _, elem := range numbers {
		if elem < min {
			min = elem
		}
		if elem > max {
			max = elem
		}
	}
	return min, max
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resMin, resMax := minMaxNum(numbers)
	fmt.Println("Минимальное значение -", resMin)
	fmt.Println("Максимальное значение -", resMax)

}
