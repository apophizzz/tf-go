package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
)

func readStopWords(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)

	if(err != nil) {
		panic(fmt.Sprintf("Error while reading file: %s", filename))
	}

	return strings.Split(string(bytes), ",")
}



func readWords(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)

	if(err != nil) {
		panic(fmt.Sprintf("Error while reading file: %s", filename))
	}

	regex, err := regexp.Compile("[a-z]{2,}")

	if(err != nil) {
		panic("Error processing regex pattern.")
	}

	return regex.FindAllString(strings.ToLower(string(bytes)), -1)
}



func count(words []string, stopWords []string, wordFrequencies util.SortablePairList)  {

}

func printSorted(wordFrequencies util.SortablePairList)  {

}

func main() {
	stop_words := readStopWords("stop_words.txt")

	words := readWords("input.txt")

	wordFrequencyPairs := make(util.SortablePairList, 0)

	count(words, stop_words, wordFrequencyPairs)

	printSorted(wordFrequencyPairs)
}
