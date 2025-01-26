package main

import "fmt"

func findUnqElem(numsArr []int) string {
	numsMap := make(map[int]int)

	for _, val := range numsArr {
		numsMap[val]++
	}

	for num, count := range numsMap {
		if count == 1 {
			return fmt.Sprintf("Уникальное число: %d ", num)
		}
	}

	return fmt.Sprintf("Отсутствует уникальное число")
}

func main() {
	numsArr := []int{1, 1, 2, 3, 3, 3, 4, 4}
	res := findUnqElem(numsArr)

	fmt.Println(res)
}
