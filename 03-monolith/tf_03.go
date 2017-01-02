package main

import (
	"github.com/PaddySmalls/golang-term-frequency/util"
	"io/ioutil"
	"strings"
	"log"
	"bufio"
	"os"
)

// Allocate slice for word-frequency pairs.
var wordFrequencyPairs util.SortablePairList = make(util.SortablePairList, 0)

// Allocate slice for stop words.
var stopWords []string = make([]string, 0)



func main() {

	// Prepare stop words list.
	buffer, err := ioutil.ReadFile("stop_words.txt")

	if (err != nil) {
		panic(err)
	}

	stopWords = strings.Split(string(buffer), ",")
	stopWords = append(stopWords, util.CreateLowercaseAlphabet()...)


	// -----------------------------------------------------------------------------

	// Open input file and prepare scanner.
	inputFile, err := os.Open("input.txt")

	if (err != nil) {
		panic(err)
	}

	inputFileScanner := bufio.NewScanner(inputFile)
	inputFileScanner.Split(bufio.ScanLines)

	// -----------------------------------------------------------------------------

	for inputFileScanner.Scan() {
		nextLine := inputFileScanner.Text()
		log.Println(nextLine)
	}

}