package main

import "fmt"

func groupPref(words []string, preflen int) map[string][]string {
	resMap := make(map[string][]string)

	for _, word := range words {
		if len(word) < preflen {
			preflen = len(word)
		}
		prefix := word[0:preflen]
		resMap[prefix] = append(resMap[prefix], word)
	}

	return resMap
}

func main() {
	words := []string{"apple", "appreciate", "bat", "ball", "ban", "banana", "cat", "catalog"}
	res := groupPref(words, 2)

	fmt.Println(res)
}
