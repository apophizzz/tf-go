package main

import (
	"io/ioutil"
	"regexp"
	"strings"
	"log"
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"sort"
)

func readFile(filePath string) string {
	fileData, err := ioutil.ReadFile(filePath)

	if (err != nil) {
		panic(err)
	}

	return string(fileData)
}

func filterCharsAndNormalize(data string) string {
	regexp, err := regexp.Compile("[\\W_]+")

	if (err != nil) {
		panic(err)
	}

	return strings.ToLower(regexp.ReplaceAllLiteralString(data, " "))
}

func scan(data string) []string {
	return strings.Split(data, " ")
}

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

func sortWfPairs(wfPairs util.SortablePairList) util.SortablePairList {
	pairList := wfPairs
	sort.Sort(pairList)
	return pairList
}

func printAll(wfPairs util.SortablePairList) {
	if (len(wfPairs) > 0) {
		log.Printf("Word: %s - Frequency: %d", wfPairs[0].Key, wfPairs[0].Val)
		printAll(wfPairs[1:])
	}
}

func main() {
	printAll(sortWfPairs(countFrequencies(removeStopWords(scan(filterCharsAndNormalize(readFile("input.txt")))))))
}
