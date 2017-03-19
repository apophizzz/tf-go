package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"io/ioutil"
	"strings"
	"log"
	"bufio"
	"os"
	"sort"
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

		var startCharIndex int = -1

		// Iterate over line's characters
		for index, char := range nextLine {
			if (startCharIndex == -1) {
				if (util.IsAlphanumeric(char)) {
					// Start of a new word has been found.
					startCharIndex = index
				}
			} else {
				// If startCharIndex has already been set ...
				if (!util.IsAlphanumeric(char)) {
					// End of a word has been found.
					wordFound := nextLine[startCharIndex:index]

					// Verify that word is not a stop word.
					if (!util.Contains(stopWords, wordFound)) {
						var wordExists bool

						// Loop over all available word-frequency pairs.
						for _, wfPair := range wordFrequencyPairs {
							if (wordFound == wfPair.Key) {
								// Word is already present in pair list, increment count and leave loop.
								wfPair.Val++
								wordExists = true
								break
							}
						}

						if (!wordExists) {
							// Word is not present in list yet. Add it.
							wordFrequencyPairs = append(wordFrequencyPairs, &util.SortablePair{Key:wordFound, Val:1})
						}
					}

					// Reset startCharIndex for being able to search for the next word.
					startCharIndex = -1
				}
			}
		}
	}

	// -----------------------------------------------------------------------------

	// All lines have been processed, sort word-frequency pair list.
	sort.Sort(wordFrequencyPairs)

	// Print top 25 words.
	for index, wfPair := range wordFrequencyPairs {
		if(index < 25) {
			log.Printf("Word: %s - Frequency: %d\n", wfPair.Key, wfPair.Val)
		} else {
			break
		}
	}
}