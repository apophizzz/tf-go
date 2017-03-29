package filter

import (
	"io/ioutil"
	"fmt"
	"strings"
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
)

type StopWordsManager struct {
	stopWords []string
}

func (swm *StopWordsManager) Init(fileName string) {
	stopWordsFileContent := readFromStopWordsFile(fileName)
	swm.stopWords = append(strings.Split(stopWordsFileContent, ","), util.CreateLowercaseAlphabet()...)
}

func readFromStopWordsFile(fileName string) string {
	bytes , err := ioutil.ReadFile(fileName)

	if(err != nil) {
		panic(fmt.Sprintf("Something went wrong while reading stop words from file: %s", fileName))
	}

	return string(bytes)
}

/*
	Only for debugging purposes.
 */
func (swm *StopWordsManager) Show() {
	fmt.Println(swm.stopWords)
}


