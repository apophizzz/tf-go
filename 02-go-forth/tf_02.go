package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
	"sort"
)

// Prepare stack.
var stack util.Stack = util.Stack{}

// Prepare heap.
var heap map[string]interface{} = make(map[string]interface{})


/*
	STEP 1:
	Read the entire input file and store the whole content as a single string.
	This string is then pushed onto the stack.
 */
func readInputFile() {
	heap["buffer"], heap["err"] = ioutil.ReadFile("input.txt")

	if (heap["err"] != nil) {
		panic(heap["err"])
	}

	stack.Push(util.StackElement{Val:string(heap["buffer"].([]byte))})


	// Clear heap.
	delete(heap, "buffer")
	delete(heap, "err")
}

/*
	STEP 2:
	Create regex pattern on heap (due to multi-value return), apply it to the input (which has been pushed
	onto the stack by "readInputFile"), and push the cleaned string back to the stack.
 */
func filterInvalidChars() {
	heap["regexp"], heap["err"] = regexp.Compile("[\\W_]+")

	if (heap["err"] != nil) {
		panic(heap["err"])
	}

	heap["input"], heap["err"] = stack.Pop()
	heap["input"] = heap["input"].(util.StackElement).Val


	stack.Push(util.StackElement{Val:heap["regexp"].(*regexp.Regexp).
		ReplaceAllLiteralString(heap["input"].(string), " ")})

	// Clear heap.
	delete(heap, "err")
	delete(heap, "regexp")
	delete(heap, "input")
}

/*
	STEP 3:
	Grab the clean input string from the stack, split it into single words and push these words
	to the stack.
 */
func scan() {
	heap["input"], heap["err"] = stack.Pop()
	heap["input"] = heap["input"].(util.StackElement).Val
	heap["input"] = strings.Split(heap["input"].(string), " ")

	for _, word := range heap["input"].([]string) {
		stack.Push(util.StackElement{Val:word})
	}

	// Clear heap.
	delete(heap, "input")
	delete(heap, "err")
}


/*
	STEP 4:
	Read stop words file, extend it expand it by the lowercase ASCII letters and store the resulting array on the
	heap. Then, pop the input from the stack word by word and store it in case it's not a stop word. Otherwise,
	it simply gets discarded. At the end, the array of valid words is pushed onto the stack for further processing.
 */
func removeStopWords() {
	heap["stopWordsBytes"], heap["err"] = ioutil.ReadFile("stop_words.txt")

	if (heap["err"] != nil) {
		panic(heap["err"])
	}

	createLowercaseAlphabet()

	heap["alphabet"], _ = stack.Pop()

	heap["stopWords"] = append(strings.Split(string(heap["stopWordsBytes"].([]byte)), ","),
		heap["alphabet"].(util.StackElement).Val.([]string)...)

	heap["words"] = make([]string, 0)

	for stack.HasMoreElements() {

		heap["stackElement"], heap["err"] = stack.Pop()

		if (heap["err"] != nil) {
			panic(heap["err"])
		}

		if !util.Contains(heap["stopWords"].([]string),
			heap["stackElement"].(util.StackElement).Val.(string)) {

			heap["words"] = append(heap["words"].([]string),
				heap["stackElement"].(util.StackElement).Val.(string))
		}
	}

	stack.Push(util.StackElement{Val: heap["words"]})

	// Clear heap.
	delete(heap, "err")
	delete(heap, "stopWordsBytes")
	delete(heap, "result")
	delete(heap, "stackElement")
	delete(heap, "words")
	delete(heap, "stopWords")
	delete(heap, "alphabet")
}

func createLowercaseAlphabet() {
	heap["alphabet"] = make([]string, 0)
	heap["byteSlice"] = make([]byte, 26)

	for i := range heap["byteSlice"].([]byte) {
		heap["byteSlice"].([]byte)[i] = 'a' + byte(i)
		heap["alphabet"] = append(heap["alphabet"].([]string), string(heap["byteSlice"].([]byte)[i]))
	}

	stack.Push(util.StackElement{Val:heap["alphabet"]})

	// Clear heap.
	delete(heap, "alphabet")
	delete(heap, "byteSlice")
}

/*
	STEP 5:
	Pop the filtered slice of words from the stack. Allocate space on the heap for a map storing the frequencies,
	iterate over the word slice and update the word counts in the map. At least, push the map on the stack and
	free the heap.
 */
func computeFrequencies() {
	heap["words"], heap["err"] = stack.Pop()
	heap["words"] = heap["words"].(util.StackElement).Val

	heap["frequencies"] = make(map[string]int)

	if (heap["err"] != nil) {
		panic(heap["err"])
	}

	for heap["index"], heap["currentWord"] = range heap["words"].([]string) {

		if heap["currentWordCount"], heap["isPresent"] =
			heap["frequencies"].(map[string]int)[heap["currentWord"].(string)]; heap["isPresent"].(bool) {

			// Push current word count to stack.
			stack.Push(util.StackElement{Val:heap["currentWordCount"].(int)})

			// Push '1' to stack in order to be able to add it to the current count.
			stack.Push(util.StackElement{Val:1})

			// Pop current count and adder from stack.
			heap["adder"], heap["err"] = stack.Pop()
			heap["adder"] = heap["adder"].(util.StackElement).Val

			if (heap["err"] != nil) {
				panic(heap["err"])
			}

			heap["currentWordCount"], heap["err"] = stack.Pop()
			heap["currentWordCount"] = heap["currentWordCount"].(util.StackElement).Val

			if (heap["err"] != nil) {
				panic(heap["err"])
			}

			// Push new count to stack.
			stack.Push(util.StackElement{Val:heap["currentWordCount"].(int) + heap["adder"].(int)})

		} else {
			// Word is not present in map yet, push '1' as current count to stack.
			stack.Push(util.StackElement{Val:1})
		}

		// Pop count from stack and update count for word on heap.
		heap["newCount"], heap["err"] = stack.Pop()
		heap["newCount"] = heap["newCount"].(util.StackElement).Val

		if (heap["err"] != nil) {
			panic(heap["err"])
		}

		heap["frequencies"].(map[string]int)[heap["currentWord"].(string)] = heap["newCount"].(int)

	}

	// Push map to the stack.
	stack.Push(util.StackElement{Val:heap["frequencies"]})

	// Clear heap.
	delete(heap, "words")
	delete(heap, "err")
	delete(heap, "index")
	delete(heap, "currentWord")
	delete(heap, "currentWordCount")
	delete(heap, "isPresent")
	delete(heap, "newCount")
	delete(heap, "adder")
	delete(heap, "frequencies")
}


/*
	STEP 6:
	Fetch the unsorted map from stack and convert it to a list of sortable key-value pairs.
	Then, sort the list and push the result back onto the stack.
 */
func mapToSortedPairList() {
	heap["map"], heap["err"] = stack.Pop()
	heap["map"] = heap["map"].(util.StackElement).Val

	if (heap["err"] != nil) {
		panic(heap["err"])
	}

	// Allocate space for pair list on heap.
	heap["pairList"] = make(util.SortablePairList, 0)

	for key, val := range heap["map"].(map[string]int) {
		heap["pairList"] = append(heap["pairList"].(util.SortablePairList), &util.SortablePair{Key:key, Val:val})
	}

	sort.Sort(heap["pairList"].(util.SortablePairList))

	for i := len(heap["pairList"].(util.SortablePairList)) - 1; i >= 0; i-- {
		stack.Push(util.StackElement{ Val: heap["pairList"].(util.SortablePairList)[i]})
	}


	// Clear heap.
	delete(heap, "err")
	delete(heap, "map")
	delete(heap, "pairList")
}


/*
	STEP 7:
	Pop 25 uppermost word-frequency pairs (those with highest count) from stack and print them to console.

 */
func prettyPrintList() {

	// Start counting from zero.
	stack.Push(util.StackElement{Val:0})

	const maxPairCount = 25

	for len(stack.Elements) > 1 && stack.ElementAt(len(stack.Elements) - 1).Val.(int) < maxPairCount {

		// Pop counter from stack.
		heap["counter"], heap["err"] = stack.Pop()

		if (heap["err"] != nil) {
			panic(heap["err"])
		}

		heap["counter"] = heap["counter"].(util.StackElement).Val

		// Pop next word-frequency pair from stack.
		heap["currentPair"], heap["err"] = stack.Pop()

		if (heap["err"] != nil) {
			panic(heap["err"])
		}

		heap["currentPair"] = heap["currentPair"].(util.StackElement).Val

		log.Printf("Word: %-10s -\tFrequency: %-5d", heap["currentPair"].(*util.SortablePair).Key,
			heap["currentPair"].(*util.SortablePair).Val)

		// Push current count and adder on stack.
		stack.Push(util.StackElement{Val: heap["counter"]})
		stack.Push(util.StackElement{Val:1})

		heap["adder"], heap["err"] = stack.Pop()

		if (heap["err"] != nil) {
			panic(heap["err"])
		}

		heap["adder"] = heap["adder"].(util.StackElement).Val

		heap["counterVal"], heap["err"] = stack.Pop()

		if (heap["err"] != nil) {
			panic(heap["err"])
		}

		heap["counterVal"] = heap["counterVal"].(util.StackElement).Val

		// Push incremented counter back on stack.
		stack.Push(util.StackElement{Val:heap["counterVal"].(int) + heap["adder"].(int)})
	}

	delete(heap, "counter")
	delete(heap, "currentPair")
	delete(heap, "adder")
	delete(heap, "counterVal")
	delete(heap, "err")
}

func main() {
	readInputFile()
	filterInvalidChars()
	scan()
	removeStopWords()
	computeFrequencies()
	mapToSortedPairList()
	prettyPrintList()
}

