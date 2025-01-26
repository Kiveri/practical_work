package main

import "fmt"

func wordFreqAnal(text string) map[string]int {
	wordCount := make(map[string]int)
	wordBuild := make([]rune, 0)

	for _, char := range text {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			if char >= 'A' && char <= 'Z' {
				char += 'a' - 'A'
			}
			wordBuild = append(wordBuild, char)
		} else {
			if len(wordBuild) > 0 {
				wordCount[string(wordBuild)]++
				wordBuild = nil
			}
		}
	}

	if len(wordBuild) > 0 {
		wordCount[string(wordBuild)]++
	}

	return wordCount
}

func main() {
	text := "The game of soccer is played with a round ball. The ball is usually made of leather."
	res := wordFreqAnal(text)

	fmt.Println(res)
}
