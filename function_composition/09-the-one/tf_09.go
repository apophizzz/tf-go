package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"bytes"
)

type TFTheOne struct {
	value interface{}
}

func (theOne *TFTheOne) bind(f func(arg interface{}) interface{}) *TFTheOne {
	theOne.value = f(theOne.value)
	return theOne
}

func (theOne *TFTheOne) printMe() {
	fmt.Println(theOne.value)
}



/*
	#### FUNCTIONS ####
 */


var readFile = func(path interface{}) interface{} {
	fileBytes, err := ioutil.ReadFile(path.(string))

	if (err != nil) {
		panic(fmt.Sprintf("Unable to read file: %s", path))
	}

	return string(fileBytes)
}

var filterChars = func(fileContent interface{}) interface{} {
	regex, err := regexp.Compile("[\\W_]+")

	if (err != nil) {
		panic("An error occurred parsing regular expression.")
	}

	return regex.ReplaceAllLiteralString(fileContent.(string), " ")
}

var normalize = func(filteredContent interface{}) interface{} {
	return strings.ToLower(filteredContent.(string))
}

var scan = func(normalizedContent interface{}) interface{} {
	return strings.Fields(normalizedContent.(string))
}

var removeStopWords = func(wordList interface{}) interface{} {
	fileBytes, err := ioutil.ReadFile("stop_words.txt")

	if (err != nil) {
		panic("Unable to read file: stop_words.txt")
	}

	stopWordsList := strings.Split(string(fileBytes), ",")
	filteredWordsList := make([]string, 0)

	for _, word := range wordList.([]string) {
		if (!util.Contains(stopWordsList, word)) {
			filteredWordsList = append(filteredWordsList, word)
		}
	}

	return filteredWordsList
}

var frequencies = func(filteredWordList interface{}) interface{} {
	wfPairs := make(util.SortablePairList, 0)

	for _, word := range filteredWordList.([]string) {
		if (!wfPairs.ContainsKey(word)) {
			wfPairs.Add(&util.SortablePair{Key:word, Val:1})
		} else {
			pair, _ := wfPairs.Get(word)
			pair.Val++
		}
	}

	return &wfPairs
}

var sort = func(wfPairList interface{}) interface{} {
	pairList := wfPairList.(*util.SortablePairList)
	pairList.Sort()
	return pairList

}

var top25Frequencies = func(sortedPairs interface{}) interface{} {
	buffer := bytes.Buffer{}
	pairList := *(sortedPairs.(*util.SortablePairList))

	for i := 0; i < 25 && i < len(pairList); i++ {
		pair := pairList[i]
		buffer.WriteString(fmt.Sprintf("Word: %-10s -\tFrequency: %d\n", pair.Key, pair.Val))
	}

	return buffer.String()
}

func main() {
	theOne := TFTheOne{value:"input.txt"}
	theOne.bind(readFile).
		bind(filterChars).
		bind(normalize).
		bind(scan).
		bind(removeStopWords).
		bind(frequencies).
		bind(sort).
		bind(top25Frequencies).
		printMe()
}
