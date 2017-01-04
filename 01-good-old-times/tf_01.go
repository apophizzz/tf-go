package main

import (
	"os"
	"bufio"
	"strings"
	"unicode"
	"log"
	"strconv"
	"fmt"
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
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
	inputBytesReadCount  int
	wordStartCharIndex   int      // The start index of the next word in the current line.
	foundExistingWord    bool     // Flag indicating if word found already exists in word_freq file.
	currentWord          string   // The last word found.
	wordFrequencyPair    string   // Is "currentWord, count(currentWord)".
	currentWordFrequency int      // Frequency of last word found.
	wordFreqLineLength   int

	top25Words           []string
}


// Open input file
var inputFile *os.File = open("input.txt")

// Open secondary memory
var wordFreqsFile *os.File = touchopen("word_freqs.txt")

// Prepare primary memory for further processing
var primaryMemory FileProcessingMemory = FileProcessingMemory{}


/*

  PART 1:
  Read the input file line by line, filter stop words and write the word counts for valid non-stop words to
  "word_freqs.txt".

 */
func doPartOne() {

	// Read and store stop words
	primaryMemory.stopWords = readStopWords()

	// Read a single line from the input file. Continue until there're no more lines to be read.
	for {

		primaryMemory.currentLine, primaryMemory.inputBytesReadCount = util.ReadLineFromFile(inputFile)

		if primaryMemory.inputBytesReadCount == 0 {
			break
		}

		if (!strings.HasSuffix(primaryMemory.currentLine, "\n")) {
			primaryMemory.currentLine += "\n"
		}

		log.Printf("Current line is: '%s'", primaryMemory.currentLine)

		primaryMemory.foundExistingWord = false // Start with 'wordFound' set to false.
		primaryMemory.wordStartCharIndex = -1 // Set 'wordStartIndex' to -1


		// Loop over the current line's characters
		for index, char := range primaryMemory.currentLine {

			if (primaryMemory.wordStartCharIndex == -1) {
				if (isAlphanumeric(char)) {
					// Start of new word has been found.
					primaryMemory.wordStartCharIndex = index
				}

			} else {
				if (!isAlphanumeric(char)) {
					// End of a word has been found.
					primaryMemory.foundExistingWord = false

					primaryMemory.currentWord = strings.ToLower(
						primaryMemory.currentLine[primaryMemory.wordStartCharIndex:index])
					log.Printf("Found input word: '%s'", primaryMemory.currentWord)


					// Check if we have a valid word ..
					if (len(primaryMemory.currentWord) > 2 && !util.Contains(
						primaryMemory.stopWords, primaryMemory.currentWord)) {
						log.Printf("Check if '%s' is in stop words ...", primaryMemory.currentWord)

						// Check if word already exists in word_freq file.
						for {
							wordFreqsLine, wordFreqsBytesRead := util.ReadLineFromFile(wordFreqsFile)

							pos, _ := wordFreqsFile.Seek(0, 1)
							log.Printf("Position in word_freq file is: %d", pos)

							if (wordFreqsBytesRead == 0) {
								// End of word_freqs file reached. Leave loop.
								break
							}

							primaryMemory.wordFrequencyPair = strings.TrimSpace(string(wordFreqsLine))
							primaryMemory.wordFreqLineLength = len(primaryMemory.wordFrequencyPair)

							count, err := strconv.Atoi(strings.TrimSpace(
								strings.Split(primaryMemory.wordFrequencyPair, ",")[1]))

							if (err != nil) {
								panic(err)
							}

							primaryMemory.currentWordFrequency = count // Store count of word from target file.
							primaryMemory.wordFrequencyPair = strings.Split(
								primaryMemory.wordFrequencyPair, ",")[0] // Store current word from target file.

							// Check if current word from word_freq file equals last word found.
							if (primaryMemory.wordFrequencyPair == primaryMemory.currentWord) {
								log.Printf("Word '%s' already exists in target file!", primaryMemory.currentWord)
								primaryMemory.currentWordFrequency += 1 // Increment count by one.
								primaryMemory.foundExistingWord = true // We found an existing word.
								break
							}
						}

						if (!primaryMemory.foundExistingWord) {
							// Word found does not exist in file yet, append new entry.
							wordFreqsFile.WriteString(fmt.Sprintf("%20s,%04d\n", primaryMemory.currentWord, 1))
						} else {
							// Word found already exists, update count.
							wordFreqsFile.Seek(-26, 1)
							wordFreqsFile.WriteString(fmt.Sprintf(
								"%20s,%04d\n", primaryMemory.wordFrequencyPair, primaryMemory.currentWordFrequency))
						}
						wordFreqsFile.Seek(0, 0)
					}

					// Reset first character index of next word found
					primaryMemory.wordStartCharIndex = -1
				}
			}
		}

	}

	// Input file is no longer needed, close it properly.
	inputFile.Close()
}

/*

   PART 2: We throw away everything contained in our primary memory (no longer needed). Then our goal is to
   find the 25 most frequent words in the "word_freqs.txt" file.

 */
func doPartTwo() {
	// Clear primary memory be creating new struct
	primaryMemory = FileProcessingMemory{}

	// Prepare memory for top 25 words
	primaryMemory.top25Words = make([]string, 26, 26)

	for {
		line, byteCount := util.ReadLineFromFile(wordFreqsFile)

		if (byteCount == 0) {
			break
		}

		primaryMemory.wordFrequencyPair = strings.TrimSpace(line)
		primaryMemory.currentWord = strings.Split(primaryMemory.wordFrequencyPair, ",")[0]

		currentWordFrequency, _ := strconv.Atoi(strings.Split(primaryMemory.wordFrequencyPair, ",")[1])
		primaryMemory.currentWordFrequency = currentWordFrequency

		for i := 0; i < 25; i++ {

			var currentListWordCount int

			if (primaryMemory.top25Words[i] != "") {
				currentListWordCount, _ = strconv.Atoi(strings.Split(primaryMemory.top25Words[i], ",")[1])
			}

			if (primaryMemory.top25Words[i] == "" || currentListWordCount < primaryMemory.currentWordFrequency) {
				primaryMemory.top25Words = append(primaryMemory.top25Words[:i],
					append([]string{primaryMemory.wordFrequencyPair},
						primaryMemory.top25Words[i:]...)...)
				cropTop25Array()
				break
			}
		}

	}

	log.Println()
	log.Println("Start printing top 25 words ...")
	for _, tfEntry := range primaryMemory.top25Words {
		if (tfEntry != "") {
			log.Print(tfEntry)
		}
	}

}

func cropTop25Array() {
	arr := make([]string, 25, 25)

	for i := 0; i < 25; i++ {
		arr[i] = primaryMemory.top25Words[i]
	}
	primaryMemory.top25Words = arr
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}


func main() {
	log.Print("Start term frequency program ...\n")
	doPartOne()
	doPartTwo()
	log.Print("Finished!")
}
