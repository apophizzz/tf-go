package util

import "os"

func ReadLineFromFile(file *os.File) (string, int) {
	var line string
	var totalBytesRead int
	var currentByteAsString string

	for currentByteAsString != "\n" {
		buffer := make([]byte, 1)
		numBytesRead, _ := file.Read(buffer)

		if(numBytesRead == 0) {
			break
		}

		currentByteAsString = string(buffer)
		line += currentByteAsString
		totalBytesRead += numBytesRead
	}

	return line, totalBytesRead
}
