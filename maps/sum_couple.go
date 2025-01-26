package main

import "fmt"

func findSumUniqCoup(list map[int]int) map[int][][]int {
	sumCount := make(map[int]int)
	resMap := make(map[int][][]int)

	for k, v := range list {
		sum := k + v
		sumCount[sum]++
		resMap[sum] = append(resMap[sum], []int{k, v})
	}

	for sum, count := range sumCount {
		if count > 1 {
			delete(resMap, sum)
		}
	}

	return resMap
}

func main() {
	list := map[int]int{
		1: 6,
		2: 3,
		3: 2,
		4: 3,
		5: 1,
		6: 3,
		7: 1,
	}
	res := findSumUniqCoup(list)

	fmt.Println(res)
}
