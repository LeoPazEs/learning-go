package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsCounter := make(map[string]int)
	for _, word := range words {
		_, ok := wordsCounter[word]
		if ok {
			wordsCounter[word] += 1
		} else {
			wordsCounter[word] = 1
		}
	}
	return wordsCounter
}

func main() {
	wc.Test(WordCount)
}
