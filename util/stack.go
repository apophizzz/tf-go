package util

import (
	"errors"
)

type StackElement struct {
	Val interface{}
}


type Stack struct {
	Elements []StackElement
}

func (s *Stack) Push(elem StackElement) {
	s.Elements = append(s.Elements, elem)
}

func (s *Stack) Pop() (StackElement, error) {
	var poppedElem StackElement

	if(len(s.Elements) > 0) {
			stackLength := len(s.Elements)
			poppedElem = s.Elements[stackLength-1]
			s.Elements = s.Elements[:stackLength-1]
			return poppedElem, nil
	}

	return poppedElem, errors.New("Can't pop element from empty stack!")
}

func (s *Stack) HasMoreElements() bool {
	return len(s.Elements) > 0
}

func (s *Stack) ElementAt(index int) StackElement {
	if(index < 0 || index > (len(s.Elements)) - 1) {
		panic("Stack: Index out of bounds!")
	}

	return s.Elements[index]
}




