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


/*
	This function does nothing but end the chain of continuations
	by performing a return.
 */
var doNothing No_op = func() {
	fmt.Println("Bye, I'm outta here ...")
	return
}


/*
	Iterate over the sorted list of word-frequency pairs and print
	all the elements. Afterwards, we simply call the No_op continuation.
 */
var printAll Printer = func(pairList util.SortablePairList, cont No_op) {
	pairList.Foreach(func(pair *util.SortablePair) {
		fmt.Printf("Word: %-10s -\tFrequency: %d\n", pair.Key, pair.Val)
	})
	cont()
}


/*
	This function takes a list if word-frequency pairs as an argument and sorts
	it by word count in descending order. Afterwards, it calls the Printer
	continuation, passing the following arguments:

	1) The sorted list of word-frequency pairs.
	2) The No_op continuation to be called by the Printer after having printed
	the list of pairs.
 */
var sort Sorter = func(pairList util.SortablePairList, cont Printer) {
	pairList.Sort()
	cont(pairList, doNothing)
}


/*
	This function takes a slice containing all words which aren't stop words
	and counts them, building a list of "word-frequency pairs". After that,
	it calls the Sorter continuation it got passed, along with the following
	arguments:

	1) The list of word-frequency pairs.
	2) The function (print) to be invoked by the Sorter after it has finished
	its job.
 */
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
	cont(pairList, printAll)
}


/*
	This function takes a slice holding all the input words and discards all words
	which are considered "stop words". Afterwards, it calls the FrequencyCounter
	continuation, passing two arguments:

	1) A slice containing all words which aren't considered stop words.
	2) Another function (sort) to be invoked after the FrequencyCounter has finished
	its work.
 */
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


/*
	This function takes the input string, which is free from non-alphanumeric
	characters and uppercase letters so far, and splits it up into an array
	of words. Afterwards, it calls the StopWordsFilter continuation, passing the
	following arguments:

	1) The array containing the input string's single words.
	2) Another continuation (countFrequencies) to be called by the StopWordsFilter
	as soon as it has finished its work.
 */
var scan Scanner = func(input string, cont StopWordsFilter) {
	cont(strings.Split(input, " "), countFrequencies)
}


/*
	This function converts the input string to a lowercase string.
	Afterwards, it calls its Scanner continuation with the following arguments:

	1) The input string, converted to lowercase.
	2) The function that should be invoked by the Scanner (filterStopWords function).
 */
var normalize Normalizer = func(input string, cont Scanner) {
	cont(strings.ToLower(input), filterStopWords)
}


/*
	This function captures all non-alphanumeric characters and replaces them with
	whitespace. Afterwards, it calls its Normalizer continuation passing the
	following arguments:

	1) The purged content of the input file.
	2) The function the Normalizer should invoke after it has finished (scan function).
 */
var filterChars CharFilter = func(input string, cont Normalizer) {
	regex, err := regexp.Compile("[\\W_]+")

	if (err != nil) {
		panic("An error occurred parsing regular expression.")
	}

	cont(regex.ReplaceAllLiteralString(input, " "), scan)
}


/*
	This is where everything starts. At first, we read the whole input file's contents.
	Afterwards, we call the CharFilter continuation and pass it two arguments:

	1) The input file's contents as a string.
	2) The function to invoke (a Normalizer) after is has finished its work.
 */
func readFile(path string, cont CharFilter) {
	bytes, err := ioutil.ReadFile(path)

	if (err != nil) {
		panic(fmt.Sprintf("Unable to read file: %s", path))
	}

	// This is an example of a tail-call: The readFile function
	// calls another function as its last action. From this point,
	// the readFile function's stack frame could safely be reused.
	cont(string(bytes), normalize)
}


/*
	Style #8: "Kick Forward"

	About this style:
	This style is, in some way, a variation of the "Pipeline" style (see "05-pipeline"). The main difference between
	the "Pipeline" style and the one we implemented here is that every function has an additional parameter.
	This parameter defines the function (aka continuation) that should be invoked once a function has finished
	its work. In our concrete example, this introduces two major advantages:

	1) Any function is somewhat decoupled from the concrete continuation function it should call at the end. What
	that means is that we've introduced custom types like "Sorter", which is defined as a function that takes two
	arguments: A list of word-frequency pairs as well as a Printer continuation. As a consequence, any function which
	expects a Sorter as an argument can be passed any function that fulfills the described interface. In this way,
	different Sorters can be implemented without its clients even see any difference.

	2) If we had a compiler/runtime enabled for optimizing tail calls (which is not the case in Golang), the current
	stack frame of a function calling its continuation could be reused, since all the calling function does is return
	after the continuation call has also returned.

 */
func main() {
	readFile("input.txt", filterChars)
}
