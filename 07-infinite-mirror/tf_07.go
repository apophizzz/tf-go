package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
)

/*
	Takes the name of the file containing the stop words as an
	argument and returns an array holding these words.
 */
func readStopWords(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)

	if (err != nil) {
		panic(fmt.Sprintf("Error while reading file: %s", filename))
	}

	return strings.Split(string(bytes), ",")
}


/*
	Takes the path to the input file as an argument and returns its contents
	as an array, including every word consisting of two or more characters.
 */
func readWords(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)

	if (err != nil) {
		panic(fmt.Sprintf("Error while reading file: %s", filename))
	}

	regex, err := regexp.Compile("[a-z]{2,}")

	if (err != nil) {
		panic("Error processing regex pattern.")
	}

	return regex.FindAllString(strings.ToLower(string(bytes)), -1)
}

/*
	Counting words in a recursive manner:
	1) If 'words' is empty, return.
	2) If 'words' is not empty, take the head of the list and process it.
	3) Invoke function recursively and pass the 'rest' of the list.
 */
func count(words []string, stopWords []string, wordFrequencies *util.SortablePairList) {
	// In case of an empty words list, we simply return.
	if (len(words) == 0) {
		return
	} else {
		// Fetch the first element (head) of the word list.
		currentWord := words[0];
		if (!util.Contains(stopWords, currentWord)) {
			if (wordFrequencies.ContainsKey(currentWord)) {
				pair, _ := wordFrequencies.Get(currentWord);
				pair.Val++
			} else {
				wordFrequencies.Add(&util.SortablePair{Key:currentWord, Val:1})
			}
		}
		// Do a recursive call, passing on the "rest" of the word list.
		count(words[1:], stopWords, wordFrequencies)
	}
}


/*
	Print list of word-frequency pairs in a recursive manner:
	1) Return if pair list is empty (ends recursion).
	2) If pair list is not empty, print the first element.
	3) Pass on the 'rest' of the list whilst invoking the function recursively.
 */
func printAll(wordFrequencies *util.SortablePairList) {
	if (len(*wordFrequencies) == 0) {
		return
	} else {
		fmt.Printf("Word: %-10s - \tFrequency: %d\n", (*wordFrequencies)[0].Key, (*wordFrequencies)[0].Val)
		rest := (*wordFrequencies)[1:]
		printAll(&rest)
	}
}

/*
	This is where everything starts.
 */
func main() {
	stop_words := readStopWords("stop_words.txt")
	words := readWords("input.txt")

	wordFrequencyPairs := make(util.SortablePairList, 0)
	count(words, stop_words, &wordFrequencyPairs)
	wordFrequencyPairs.Sort()

	printAll(&wordFrequencyPairs)
}
