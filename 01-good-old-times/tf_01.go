package main

import (
	"os"
	"bufio"
	"strings"
	"unicode"
	"log"
	"strconv"
	"fmt"
	"github.com/PaddySmalls/golang-term-frequency/util"
)

func touchopen(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	return file
}

func open(filename string) *os.File {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	return file
}

func readStopWords() []string {
	// Open stop words file.
	stopWordsFile := open("stop_words.txt")

	// Read stop words to byte buffer.
	stopWordsReader := bufio.NewReader(stopWordsFile)
	stopWordsBuffer := make([]byte, 1024)
	stopWordsReader.Read(stopWordsBuffer)

	// Finished reading, close file.
	stopWordsFile.Close()

	// Convert bytes to string, split string by "," and return resulting array.
	return strings.Split(string(stopWordsBuffer), ",")
}

type FileProcessingMemory struct {
	stopWords            []string // All stop words.
	currentLine          string   // A single line of the input file.
	inputBytesReadCount int
	wordStartCharIndex   int      // The start index of the next word in the current line.
	foundExistingWord    bool     // Flag indicating if word found already exists in word_freq file.
	currentWord          string   // The last word found.
	wordFrequencyPair    string   // Is "currentWord, count(currentWord)".
	currentWordFrequency int      // Frequency of last word found.
	wordFreqLineLength int
}


// Open input file
var inputFile *os.File = open("input.txt")

// Open secondary memory
var wordFreqsFile *os.File = touchopen("word_freqs.txt")


/*

  PART 1:
  Read the input file line by line, filter stop words and write the word counts for valid non-stop words to
  "word_freqs.txt".

 */
func doPartOne() {

	memory := FileProcessingMemory{}

	// Read and store stop words
	memory.stopWords = readStopWords()

	// Read a single line from the input file. Continue until there're no more lines to be read.
	for {

		memory.currentLine, memory.inputBytesReadCount = util.ReadLineFromFile(inputFile)

		if memory.inputBytesReadCount == 0 {
			break
		}

		if (!strings.HasSuffix(memory.currentLine, "\n")) {
			memory.currentLine += "\n"
		}

		log.Printf("Current line is: '%s'", memory.currentLine)

		memory.foundExistingWord = false // Start with 'wordFound' set to false.
		memory.wordStartCharIndex = -1 // Set 'wordStartIndex' to -1


		// Loop over the current line's characters
		for index, char := range memory.currentLine {

			if (memory.wordStartCharIndex == -1) {
				if (isAlphanumeric(char)) {
					// Start of new word has been found.
					memory.wordStartCharIndex = index
				}

			} else {
				if (!isAlphanumeric(char)) {
					// End of a word has been found.
					memory.foundExistingWord = false

					memory.currentWord = strings.ToLower(memory.currentLine[memory.wordStartCharIndex:index])
					log.Printf("Found input word: '%s'", memory.currentWord)


					// Check if we have a valid word ..
					if (len(memory.currentWord) > 2 && !contains(memory.stopWords, memory.currentWord)) {
						log.Printf("Check if '%s' is in stop words ...", memory.currentWord)

						// Check if word already exists in word_freq file.
						for {
							wordFreqsLine, wordFreqsBytesRead := util.ReadLineFromFile(wordFreqsFile)

							pos,_ := wordFreqsFile.Seek(0, 1)
							log.Printf("Position in word_freq file is: %d", pos)

							if (wordFreqsBytesRead == 0) {
								// End of word_freqs file reached. Leave loop.
								break
							}

							memory.wordFrequencyPair = strings.TrimSpace(string(wordFreqsLine))
							memory.wordFreqLineLength = len(memory.wordFrequencyPair)


							count, err := strconv.Atoi(strings.TrimSpace(strings.Split(memory.wordFrequencyPair, ",")[1]))

							if (err != nil) {
								panic(err)
							}

							memory.currentWordFrequency = count // Store count of word from target file.
							memory.wordFrequencyPair = strings.Split(memory.wordFrequencyPair, ",")[0] // Store current word from target file.

							// Check if current word from word_freq file equals last word found.
							if (memory.wordFrequencyPair == memory.currentWord) {
								log.Printf("Word '%s' already exists in target file!", memory.currentWord)
								memory.currentWordFrequency += 1 // Increment count by one.
								memory.foundExistingWord = true // We found an existing word.
								break
							}
						}

						if (!memory.foundExistingWord) {
							// Word found does not exist in file yet, append new entry.
							wordFreqsFile.WriteString(fmt.Sprintf("%20s,%04d\n", memory.currentWord, 1))
						} else {
							// Word found already exists, update count.
							posOld,_ := wordFreqsFile.Seek(0, 1)
							log.Printf("Position in file before seek is: %d", posOld)
							posNew,_ := wordFreqsFile.Seek(-26, 1)
							log.Printf("Position in file after seek is: %d", posNew)

							wordFreqsFile.WriteString(fmt.Sprintf("%20s,%04d\n", memory.wordFrequencyPair, memory.currentWordFrequency))
						}
						wordFreqsFile.Seek(0, 0)
					}

					// Reset first character index of next word found
					memory.wordStartCharIndex = -1
				}
			}
		}


	}

	// Input file is no longer needed, close it properly.
	inputFile.Close()
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func contains(array []string, a string) bool {
	for _, s := range array {
		if (s == a) {
			return true
		}
	}
	return false
}

func main() {
	log.Print("Start term frequency program ...\n")
	doPartOne()
	log.Print("Finished!")
}
