package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sentence := make(map[string]int)

	for _, word := range strings.Fields(s) {
		sentence[word]++
	}
	return sentence
}

func main() {
	wc.Test(WordCount)
}
