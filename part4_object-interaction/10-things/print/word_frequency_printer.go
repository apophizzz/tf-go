package print

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/util"
	"fmt"
)

type WordFrequencyPrinter struct {
	limit int
}

func (wfp *WordFrequencyPrinter) WithLimit(limit int) {
	if(limit >= 0) {
		wfp.limit = limit
	} else {
		limit = 0
	}
}

func (wfp *WordFrequencyPrinter) PrintResult(pairs *util.SortablePairList) {
	pairList := *pairs
	for i := 0; wfp.isWithinRange(i, pairs.Len()); i++ {
		currentPair := pairList[i]
		fmt.Printf("Word: %-10s -\tFrequency: %d\n", currentPair.Key, currentPair.Val)
	}
}

func (wfp *WordFrequencyPrinter) isWithinRange(counter int, listSize int) bool {
	return counter < wfp.limit && counter < listSize
}


