package input

import (
	"strings"
	"io/ioutil"
	"fmt"
	"regexp"
)

type DataStorageManager struct {
	fileContent string
}

func (dsm *DataStorageManager) Init(fileName string) {
	lines := readLines(fileName)
	filteredLines := filterChars(lines)
	dsm.fileContent = strings.ToLower(filteredLines)
}

func readLines(fileName string) string {
	bytes , err := ioutil.ReadFile(fileName)

	if(err != nil) {
		panic(fmt.Sprintf("Something went wrong file reading file: %s", fileName))
	}

	return string(bytes)
}

func filterChars(fileContent string) string {
	pattern := "[\\W_]+";
	regex, err := regexp.Compile(pattern)

	if(err != nil) {
		panic(fmt.Sprintf("Something went wrong while parsing regular expression: %s", pattern))
	}

	return regex.ReplaceAllLiteralString(fileContent, " ")
}


func (dsm *DataStorageManager) Words() []string {
	return strings.Fields(dsm.fileContent)
}

/*
	Only for debugging purposes.
 */
func (dsm *DataStorageManager) Show()  {
	fmt.Println(dsm.fileContent)
}




