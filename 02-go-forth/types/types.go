package types

type SortablePair struct {
	Key string
	Val int
}

type SortablePairList []SortablePair


// Implement Sorter interface

func (pairList SortablePairList) Len() int {
	return len(pairList)
}

func (pairList SortablePairList) Less(i, j int) bool {
	return pairList[i].Val > pairList[j].Val
}

func (pairList SortablePairList) Swap(i, j int) {
	pairList[i], pairList[j] = pairList[j], pairList[i]
}
