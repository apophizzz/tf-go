package util

import "errors"

type StackElement struct {
	Val interface{}
}


type Stack struct {
	Elements []StackElement
}

func Push(s *Stack, elem StackElement) {
	s.Elements = append(s.Elements, elem)
}

func Pop(s *Stack) (StackElement, error) {
	var poppedElem StackElement
	if(len(s.Elements) > 0) {
			stackLength := len(s.Elements)
			poppedElem = s.Elements[stackLength-1]
			s.Elements = s.Elements[:stackLength-1]
			return poppedElem, nil
	}

	return poppedElem, errors.New("Can't pop element from empty stack!")
}

func HasMoreElements(s *Stack) bool {
	return len(s.Elements) > 0
}

func ElementAt(s* Stack, index int) StackElement {
	if(index < 0 || index > (len(s.Elements)) - 1) {
		panic("Stack: Index out of bounds!")
	}

	return s.Elements[index]
}




