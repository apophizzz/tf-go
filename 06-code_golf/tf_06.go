package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"fmt"
	"strings"
)

/*
	The "Code Golf" style aims at keeping the program performing a certain
	exercise as short as possible. This is generally reached by using
	built-in or third-party libraries and abstractions.
	Since there're not so much high-level functions for the term frequency
	task in Go (or rather I'm not familiar with the existing ones), I implemented a bunch
	of functions which perform the low-level work and therefore act as my own
	"utility functions" in this context.
 */
func main() {

	stopWords := util.Scan(util.ContentAsString("stop_words.txt"), ",")
	words := util.FindMatching(strings.ToLower(util.ContentAsString("input.txt")), "[a-z]{2,}")

	wfPairs := util.CountWords(util.Filter(words, stopWords))
	wfPairs.Sort()
	wfPairs.Foreach(func(pair *util.SortablePair) {
		fmt.Printf("Word: %-10s -\tFrequency: %d\n", pair.Key, pair.Val)
	})

}
