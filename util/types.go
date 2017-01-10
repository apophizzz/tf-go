package util

import (
	"fmt"
	"errors"
)

type SortablePair struct {
	Key string
	Val int
}

type SortablePairList []*SortablePair


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

func (pairList SortablePairList) ContainsKey(key string) bool {
	for _, pair := range pairList {
		if (key == pair.Key) {
			return true
		}
	}
	return false
}

func (pairList SortablePairList) Get(key string) (*SortablePair, error) {
	for _, pair := range pairList {
		if (key == pair.Key) {
			return pair, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("No key matching %s.", key))
}
