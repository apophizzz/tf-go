package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"io/ioutil"
	"unicode"
	"unicode/utf8"
	"strings"
	"log"
	"sort"
)


var data []byte = make([]byte, 0)
var words []string = make([]string, 0)
var wfPairs util.SortablePairList = make(util.SortablePairList, 0)


/*
	Read whole input file and store the resulting byte array
	in the global 'data' variable.
 */
func readFile(filePath string)  {
	fileData, err := ioutil.ReadFile(filePath)

	if(err != nil) {
		panic(err)
	}

	data = fileData
}


/*
	Iterate over the input file bytes and replace every
	non-alphanumeric character by whitespace.
 */
func filterCharsAndNormalize() {
	for index := range data {
		r, _ := utf8.DecodeRune(data[index:index+1])
		if(!util.IsAlphanumeric(r)) {
			utf8.EncodeRune(data[index:index+1], ' ')
		} else {
			utf8.EncodeRune(data[index:index+1], unicode.ToLower(r))
		}
	}
}


/*
	Scan input data for words and store them in the global variable 'words'.
 */
func scan() {
	words = append(words, strings.Split(string(data), " ")...)
}


/*
	Read stop words from file and filter them from the list of words.
 */
func removeStopWords() {
	stopWordsBytes, _ := ioutil.ReadFile("stop_words.txt")

	stopWords := strings.Split(string(stopWordsBytes), ",")
	stopWords = append(stopWords, util.CreateLowercaseAlphabet()...)

	for index, word := range words {
		if(util.Contains(stopWords, word)) {
			if(index < len(words)) {
				words = append(words[:index], words[index+1:]...)
			} else {
				words = words[:index]
			}
		}
	}
}


/*
	Iterate over filtered list of words and count their appearances.
 */
func computeFrequencies() {
	for _, word := range words {
		wordExists := false
		for	_, pair := range wfPairs {
			if(pair.Key == word) {
				wordExists = true
				pair.Val++
				break
			}
		}

		if(!wordExists) {
			wfPairs = append(wfPairs, &util.SortablePair{Key:word, Val:1})
		}
	}
}


/*
	Sort list of word-frequency pairs by their value (i.e. word count)
	in descending order.
 */
func sortWordFrequencyPairs() {
	sort.Sort(wfPairs)
}


/*
	Print 25 words with highest frequency to console.
 */
func printTop25() {
	for index, pair := range wfPairs {
		if(index < 25) {
			log.Printf("Word: %s - Frequency: %d\n", pair.Key, pair.Val)
		} else {
			// Stop printing results when top 25 words have been printed.
			break
		}
	}
}


func main() {
	readFile("input.txt")
	filterCharsAndNormalize()
	scan()
	removeStopWords()
	computeFrequencies()
	sortWordFrequencyPairs()
	printTop25()
}

