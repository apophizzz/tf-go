package util

import (
	"os"
	"strings"
)

func ReadLineFromFile(file *os.File) (string, int) {
	var line string
	var totalBytesRead int
	var currentByteAsString string

	for currentByteAsString != "\n" {
		buffer := make([]byte, 1)
		numBytesRead, _ := file.Read(buffer)

		if (numBytesRead == 0) {
			break
		}

		currentByteAsString = string(buffer)
		line += currentByteAsString
		totalBytesRead += numBytesRead
	}

	return line, totalBytesRead
}

func Contains(array []string, a string) bool {
	for _, s := range array {
		if (s == a) {
			return true
		}
	}
	return false
}



func CreateLowercaseAlphabet() []string {
	byteSlice := make([]byte, 26)

	for i := range byteSlice {
		byteSlice[i] = 'a' + byte(i)
	}

	return strings.Split(string(byteSlice), "")
}