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

func main() {
	readInputFile()
	filterInvalidChars()
	scan()

	log.Print(stack.Elements)
	log.Print(heap)
}

