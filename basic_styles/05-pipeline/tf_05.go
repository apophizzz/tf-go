package main

import (
	"io/ioutil"
	"regexp"
	"strings"
	"log"
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"sort"
)


/*
	Read contents of the input file specified by 'filePath'
	and return them as a single string.
 */
func readFile(filePath string) string {
	fileData, err := ioutil.ReadFile(filePath)

	if (err != nil) {
		panic(err)
	}

	return string(fileData)
}


/*
	Replace all non-alphanumeric characters within 'data' with
	whitespace and return clean string.
 */
func filterCharsAndNormalize(data string) string {
	regex, err := regexp.Compile("[\\W_]+")

	if (err != nil) {
		panic(err)
	}

	return strings.ToLower(regex.ReplaceAllLiteralString(data, " "))
}


/*
	Split 'data' string into an array of single words.
 */
func scan(data string) []string {
	return strings.Split(data, " ")
}


/*
	Remove all stop words from 'words' array.
 */
func removeStopWords(words []string) []string {
	stopWordsFileBytes, err := ioutil.ReadFile("stop_words.txt")

	if (err != nil) {
		panic(err)
	}

	stopWords := strings.Split(string(stopWordsFileBytes), ",")
	validWords := make([]string, 0)

	for _, word := range words {
		if (!util.Contains(stopWords, word)) {
			validWords = append(validWords, word)
		}
	}

	return validWords
}


/*
	Iterate through all valid words and count their appearances.
 */
func countFrequencies(words []string) util.SortablePairList {
	wfPairs := make(util.SortablePairList, 0)

	for _, word := range words {
		if wfPairs.ContainsKey(word) {
			pair, _ := wfPairs.Get(word)
			pair.Val++
		} else {
			wfPairs = append(wfPairs, &util.SortablePair{Key:word, Val:1})
		}
	}

	return wfPairs
}


/*
	Take a list of sortable pairs as an argument and return
	a sorted version of it.
 */
func sortWfPairs(wfPairs util.SortablePairList) util.SortablePairList {
	pairList := wfPairs
	sort.Sort(pairList)
	return pairList
}


/*
	Iterate through the list of sorted word-frequency pairs and print
	them one by one.
 */
func printAll(wfPairs util.SortablePairList) {
	if (len(wfPairs) > 0) {
		log.Printf("Word: %-10s -\tFrequency: %d", wfPairs[0].Key, wfPairs[0].Val)
		printAll(wfPairs[1:])
	}
}


/*
	The starting point.
 */
func main() {
	printAll(sortWfPairs(countFrequencies(removeStopWords(scan(filterCharsAndNormalize(readFile("input.txt")))))))
}
