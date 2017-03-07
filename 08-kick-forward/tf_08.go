package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"io/ioutil"
	"fmt"
	"strings"
	"regexp"
)

type No_op func()
type Printer func(util.SortablePairList, No_op)
type Sorter func(util.SortablePairList, Printer)
type FrequencyCounter func([]string, Sorter)
type StopWordsFilter func([]string, FrequencyCounter)
type Scanner func(string, StopWordsFilter)
type Normalizer func(string, Scanner)
type CharFilter func(string, Normalizer)
type FileReader func(string, CharFilter)

var doNothing No_op = func() {
	fmt.Println("Bye, I'm outta here ...")
	return
}

var print Printer = func(pairList util.SortablePairList, cont No_op) {
	pairList.Foreach(func(pair *util.SortablePair) {
		fmt.Printf("Word: %-10s -\tFrequency: %d\n", pair.Key, pair.Val)
	})
	cont()
}

var sort Sorter = func(pairList util.SortablePairList, cont Printer) {
	pairList.Sort()
	cont(pairList, doNothing)
}

var countFrequencies = func(words []string, cont Sorter) {
	pairList := make(util.SortablePairList, 0)
	for _, word := range words {
		if (pairList.ContainsKey(word)) {
			pair, _ := pairList.Get(word)
			pair.Val++
		} else {
			pairList.Add(&util.SortablePair{Key:word, Val:1})
		}
	}
	cont(pairList, print)
}

var filterStopWords StopWordsFilter = func(words []string, cont FrequencyCounter) {
	bytes, err := ioutil.ReadFile("stop_words.txt")

	if (err != nil) {
		panic("Unable to read from file: stop_words.txt")
	}

	stopWordsList := strings.Split(string(bytes), ",")
	filteredWords := make([]string, 0)

	for _, word := range words {
		if (!util.Contains(stopWordsList, word)) {
			filteredWords = append(filteredWords, word)
		}
	}

	cont(filteredWords, sort)
}

var scan Scanner = func(input string, cont StopWordsFilter) {
	cont(strings.Split(input, " "), countFrequencies)
}

var normalize Normalizer = func(input string, cont Scanner) {
	cont(strings.ToLower(input), filterStopWords)
}

var filterChars CharFilter = func(input string, cont Normalizer) {
	regex, err := regexp.Compile("[\\W_]+")

	if(err != nil) {
		panic("An error occurred parsing regular expression.")
	}

	cont(regex.ReplaceAllLiteralString(input, " "), scan)
}



func readFile(path string, cont CharFilter) {
	bytes, err := ioutil.ReadFile(path)

	if (err != nil) {
		panic(fmt.Sprintf("Unable to read file: %s", path))
	}

	cont(string(bytes), normalize)
}

func main() {
	readFile("input.txt", doNothing)
}
