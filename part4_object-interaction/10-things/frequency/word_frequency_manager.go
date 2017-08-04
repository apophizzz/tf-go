package frequency

import "github.com/apophis90/tf-go/util"

type WordFrequencyManager struct {
	wordFrequencyPairs *util.SortablePairList
}

func (wfm *WordFrequencyManager) Init() {
	allocatedList := make(util.SortablePairList, 0)
	wfm.wordFrequencyPairs = &allocatedList
}

func (wfm *WordFrequencyManager) IncrementCount(word string) {
	if wfm.wordFrequencyPairs.ContainsKey(word) {
		matchingPair, _ := wfm.wordFrequencyPairs.Get(word)
		matchingPair.IncrementCount()
	} else {
		wfm.wordFrequencyPairs.Add(&util.SortablePair{Key: word, Val: 1})
	}
}

func (wfm *WordFrequencyManager) GetSortedPairs() *util.SortablePairList {
	wfm.wordFrequencyPairs.Sort()
	return wfm.wordFrequencyPairs
}
