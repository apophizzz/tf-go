package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/part4_object-interaction/10-things/input"
	"github.com/PaddySmalls/golang_term-frequency-styles/part4_object-interaction/10-things/filter"
)

var dataStorageManager = input.DataStorageManager{}
var stopWordsManager = filter.StopWordsManager{}

func main() {
	dataStorageManager.Init("input.txt")
	dataStorageManager.Show()

	stopWordsManager.Init("stop_words.txt")
	stopWordsManager.Show()
}


