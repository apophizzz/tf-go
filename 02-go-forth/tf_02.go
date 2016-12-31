package main

import (
	"github.com/PaddySmalls/golang-term-frequency/util"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
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

	util.Push(&stack, util.StackElement{Val:string(heap["buffer"].([]byte))})

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

	heap["input"], heap["err"] = util.Pop(&stack)
	heap["input"] = heap["input"].(util.StackElement).Val

	util.Push(&stack, util.StackElement{Val:heap["regexp"].(*regexp.Regexp).
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
	heap["input"], heap["err"] = util.Pop(&stack)
	heap["input"] = heap["input"].(util.StackElement).Val
	heap["input"] = strings.Split(heap["input"].(string), " ")

	for _, word := range heap["input"].([]string) {
		util.Push(&stack, util.StackElement{Val:word})
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
	heap["buffer"], heap["err"] = ioutil.ReadFile("stop_words.txt")

	if (heap["err"] != nil) {
		panic(heap["err"])
	}

	createLowercaseAlphabet()
	heap["callResult"], _ = util.Pop(&stack)

	heap["stopWords"] = append(strings.Split(string(heap["buffer"].([]byte)), ","),
		heap["callResult"].(util.StackElement).Val.([]string)...)

	heap["words"] = make([]string, 0)

	for util.HasMoreElements(&stack) {

		heap["stackElement"], heap["err"] = util.Pop(&stack)

		if (heap["err"] != nil) {
			panic(heap["err"])
		}

		if !util.Contains(heap["stopWords"].([]string),
			heap["stackElement"].(util.StackElement).Val.(string)) {
			heap["words"] = append(heap["words"].([]string),
				heap["stackElement"].(util.StackElement).Val.(string))
		}
	}

	util.Push(&stack, util.StackElement{Val: heap["words"]})

	// Clear heap.
	delete(heap, "err")
	delete(heap, "buffer")
	delete(heap, "result")
	delete(heap, "stackElement")
	delete(heap, "words")
	delete(heap, "stopWords")
	delete(heap, "callResult")
}

func createLowercaseAlphabet() {
	heap["result"] = make([]string, 0)
	heap["byteSlice"] = make([]byte, 26)

	for i := range heap["byteSlice"].([]byte) {
		heap["byteSlice"].([]byte)[i] = 'a' + byte(i)
		heap["result"] = append(heap["result"].([]string), string(heap["byteSlice"].([]byte)[i]))
	}

	util.Push(&stack, util.StackElement{Val:heap["result"]})

	delete(heap, "result")
	delete(heap, "byteSlice")
}


func computeFrequencies() {

}

func main() {
	readInputFile()
	filterInvalidChars()
	scan()
	removeStopWords()

	log.Print(stack.Elements)
	log.Print(heap)
}

