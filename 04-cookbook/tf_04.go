package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"io/ioutil"
	"unicode"
	"unicode/utf8"
)


var data []byte = make([]byte, 0)
var words []string = make([]string, 0)
var wfPairs util.SortablePairList = make(util.SortablePairList, 0)


/*
	Read whole input file and store the resulting byte array
	in the global 'data' variable.
 */
func readFile(filePath string)  {
	fileData, err := ioutil.ReadFile(filePath)

	if(err != nil) {
		panic(err)
	}

	data = fileData
}


/*
	Iterate over the input file bytes and replace every
	non-alphanumeric character by whitespace.
 */
func filterCharsAndNormalize() {
	for index := range data {
		r, _ := utf8.DecodeRune(data[index:index+1])
		if(!util.IsAlphanumeric(r)) {
			utf8.EncodeRune(data[index:index+1], ' ')
		} else {
			utf8.EncodeRune(data[index:index+1], unicode.ToLower(r))
		}
	}
}



func main() {
	readFile("input.txt")
	filterCharsAndNormalize()
}

