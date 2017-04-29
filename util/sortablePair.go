package util

import (
	"fmt"
	"errors"
	"sort"
)

type SortablePair struct {
	Key string
	Val int
}

func (pair *SortablePair) IncrementCount() {
	pair.Val++
}

type SortablePairList []*SortablePair


// Implement Sorter interface

func (pairList *SortablePairList) Len() int {
	return len(*pairList)
}

func (pairList *SortablePairList) Less(i, j int) bool {
	list := *pairList
	return list[i].Val > list[j].Val
}

func (pairList *SortablePairList) Swap(i, j int) {
	list := *pairList
	list[i], list[j] = list[j], list[i]
}

/*
	Check if the pair list contains any pair with
	the given key. Returns true if key exists, false otherwise.
 */
func (pairList *SortablePairList) ContainsKey(key string) bool {
	for _, pair := range *pairList {
		if (key == pair.Key) {
			return true
		}
	}
	return false
}

/*
	Retrieve an element by the corresponding key. If no such key
	exists, an error is returned.
 */
func (pairList *SortablePairList) Get(key string) (*SortablePair, error) {
	list := *pairList
	for _, pair := range list {
		if (key == pair.Key) {
			return pair, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("No key matching %s.", key))
}

/*
	Appends an element to the list.
 */
func (pairList *SortablePairList) Add(pair *SortablePair)  {
	*pairList = append(*pairList, pair)
}

/*
	Executes a sorting of the list of pairs according to the
	implementation of the Sorter interface.
 */
func (pairList *SortablePairList) Sort() {
	sort.Sort(pairList)
}

/*
	Takes a function which is subsequently applied on each element
	of the list of sortable pairs.
 */
func (pairList *SortablePairList) Foreach(f func(*SortablePair))  {
	list := *pairList
	for _, pair := range list {
		f(pair)
	}
}
