package main

import "github.com/PaddySmalls/golang_term-frequency-styles/util"

func readStopWords(filename string) []string {
	return nil
}

func readWords(filename string) []string {
	return nil
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
