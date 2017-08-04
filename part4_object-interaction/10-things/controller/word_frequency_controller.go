package controller

import (
	"github.com/apophis90/tf-go/part4_object-interaction/10-things/filter"
	"github.com/apophis90/tf-go/part4_object-interaction/10-things/frequency"
	"github.com/apophis90/tf-go/part4_object-interaction/10-things/input"
	"github.com/apophis90/tf-go/part4_object-interaction/10-things/print"
)

type WordFrequencyController struct {
	dataStorageManager   *input.DataStorageManager
	stopWordsManager     *filter.StopWordsManager
	wordFrequencyManager *frequency.WordFrequencyManager
	wordFrequencyPrinter *print.WordFrequencyPrinter
}

func (wfc *WordFrequencyController) Init(inputFilePath string) {
	wfc.dataStorageManager = &input.DataStorageManager{}
	wfc.dataStorageManager.Init(inputFilePath)

	wfc.stopWordsManager = &filter.StopWordsManager{}
	wfc.stopWordsManager.Init("../../stop_words.txt")

	wfc.wordFrequencyManager = &frequency.WordFrequencyManager{}
	wfc.wordFrequencyManager.Init()

	wfc.wordFrequencyPrinter = &print.WordFrequencyPrinter{}
	wfc.wordFrequencyPrinter.WithLimit(25)

}

func (wfc *WordFrequencyController) Run() {
	inputFileWords := wfc.dataStorageManager.Words()

	for _, word := range inputFileWords {
		if !wfc.stopWordsManager.IsStopWord(word) {
			wfc.wordFrequencyManager.IncrementCount(word)
		}
	}

	wfc.wordFrequencyPrinter.PrintResult(wfc.wordFrequencyManager.GetSortedPairs())
}
