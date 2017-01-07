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

func readFile(filePath string)  {
	fileData, err := ioutil.ReadFile(filePath)

	if(err != nil) {
		panic(err)
	}

	data = fileData
}

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

