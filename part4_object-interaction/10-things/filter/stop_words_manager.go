package filter

import (
	"fmt"
	"github.com/apophis90/tf-go/util"
	"io/ioutil"
	"strings"
)

type StopWordsManager struct {
	stopWords []string
}

func (swm *StopWordsManager) Init(fileName string) {
	stopWordsFileContent := readFromStopWordsFile(fileName)
	swm.stopWords = append(strings.Split(stopWordsFileContent, ","), util.CreateLowercaseAlphabet()...)
}

func readFromStopWordsFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(fmt.Sprintf("Something went wrong while reading stop words from file: %s", fileName))
	}

	return string(bytes)
}

func (swm *StopWordsManager) IsStopWord(word string) bool {
	return util.Contains(swm.stopWords, word)
}

/*
	Only for debugging purposes.
*/
func (swm *StopWordsManager) Show() {
	fmt.Println(swm.stopWords)
}
