package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		_, ok := count[word]
		if ok {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
