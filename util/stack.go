package util

import "errors"

type StackElement struct {
	Val interface{}
}


type Stack struct {
	elements []StackElement
}

func Push(s *Stack, elem StackElement) {
	s.elements = append(s.elements, elem)
}

func Pop(s *Stack) (StackElement, error) {
	var poppedElem StackElement
	if(len(s.elements) > 0) {
			stackLength := len(s.elements)
			poppedElem = s.elements[stackLength-1]
			s.elements = s.elements[:stackLength-1]
			return poppedElem, nil
	}

	return poppedElem, errors.New("Can't pop element from empty stack!")
}




