package util

import (
	"regexp"
	"fmt"
	"strings"
)

func FindMatching(input string, regexString string) []string {
	if(len(input) == 0) {
		panic("Expected valid input string!")
	}

	if(len(regexString) == 0) {
		panic("Expected valid regex string!")
	}

	regex, err := regexp.Compile(regexString)

	if(err != nil) {
		panic(fmt.Errorf("Something went wrong while compiling regex '%s'!", regexString))
	}

	return regex.FindAllString(input, -1)

}


func Scan(input string, separator string) []string {
	if(len(input) == 0) {
		panic("Valid input string expected!")
	}

	return strings.Split(input, separator)
}


func Filter(input []string, toFilter []string) []string {
	if(input == nil || toFilter == nil) {
		panic("Valid input and list of words to filter expected!")
	}

	filteredWords := make([]string, 0)

	for _, word := range input {
		if(!Contains(toFilter, word)) {
			filteredWords = append(filteredWords, word)
		}
	}

	return filteredWords
}


func CountWords(input []string) SortablePairList {
	wordFrequencyPairs := make(SortablePairList, 0)

	for _, word := range input {
		if(wordFrequencyPairs.ContainsKey(word)) {
			pair, _ := wordFrequencyPairs.Get(word)
			pair.Val++
		} else {
			wordFrequencyPairs.Add(&SortablePair{Key:word, Val:1})
		}
	}

	return wordFrequencyPairs
}
