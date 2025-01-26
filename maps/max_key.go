package main

import "fmt"

func maxKeys(list map[string]int) []string {
	maxKeysArr := make([]string, 0)
	var maxVal int

	for k, v := range list {
		if v > maxVal {
			maxVal = v
			maxKeysArr = []string{k}
		} else if v == maxVal {
			maxKeysArr = append(maxKeysArr, k)
		}
	}

	return maxKeysArr
}

func main() {
	list := map[string]int{
		"a": 1,
		"b": 2,
		"c": 2,
		"d": 3,
		"e": 5,
		"f": 8,
		"g": 8,
		"h": 15,
		"i": 15,
		"j": 15,
	}

	res := maxKeys(list)
	fmt.Println("Max value: ", res)

}
