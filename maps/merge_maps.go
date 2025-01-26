package main

import "fmt"

func mergeMapsMaxVal(first, second map[string]int) map[string]int {
	resMap := make(map[string]int)

	for k, v := range first {
		resMap[k] = v
	}

	for k, v := range second {
		if value, exists := resMap[k]; exists {
			if v > value {
				resMap[k] = v
			}
		} else {
			resMap[k] = v
		}
	}

	return resMap
}

func main() {
	first := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	second := map[string]int{
		"c": 6,
		"d": 7,
		"e": 8,
		"f": 9,
	}
	res := mergeMapsMaxVal(first, second)

	fmt.Println(res)
}
