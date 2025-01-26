package main

import "fmt"

func evenOddNum(numbers []int) (int, int) {
	var even int
	var odd int
	for _, elem := range numbers {
		if elem%2 == 0 {
			even++
		}
		if elem%2 != 0 {
			odd++
		}
	}
	return even, odd
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resEven, resOdd := evenOddNum(numbers)
	fmt.Println("Кол-во четных -", resEven)
	fmt.Println("Кол-во нечетных - ", resOdd)

}
