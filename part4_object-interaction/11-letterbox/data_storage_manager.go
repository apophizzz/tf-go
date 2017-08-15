package storage

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type DataStorageManager struct {
	fileContent string
}

func (dsm *DataStorageManager) Dispatch(message []string) ([]string, error) {
	if !isEmptyMessage(message) {
		if isValidInitMessage(message) {
			return nil, dsm.handleInitMessage(message)
		} else if isValidWordsMessage(message) {
			return asWordsList(dsm.fileContent), nil
		}
		return nil, &InvalidDataStorageManagerMessage{message}
	} else {
		return nil, &EmptyDataStorageManagerMessage{}
	}
	return nil, nil
}

func (dsm *DataStorageManager) handleInitMessage(message []string) error {
	content, inputErr := readFileContent(message[1])

	if inputErr != nil {
		return inputErr
	}

	dsm.fileContent = strings.ToLower(
		cutTailingAndLeadingWhitespace(
			replaceNonalnumCharsByWhitespace(content)))
	return nil
}

func isEmptyMessage(message []string) bool {
	return len(message) == 0
}

func isValidInitMessage(message []string) bool {
	return (len(message) == 2 && message[0] == "init")
}

func isValidWordsMessage(message []string) bool {
	return message[0] == "words"
}

func asWordsList(fileContent string) []string {
	return strings.Fields(fileContent)
}

func readFileContent(filename string) (string, error) {
	var content string
	var fileMissing error

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fileMissing = &MissingInputFile{filename}
	} else {
		content = string(bytes)
	}
	return content, fileMissing
}

func replaceNonalnumCharsByWhitespace(fileContent string) string {
	regexPattern := "[\\W_]+"
	regex, _ := regexp.Compile(regexPattern)
	return regex.ReplaceAllLiteralString(fileContent, " ")
}

func cutTailingAndLeadingWhitespace(fileContent string) string {
	return strings.TrimSpace(fileContent)
}

type EmptyDataStorageManagerMessage struct {
}

func (err *EmptyDataStorageManagerMessage) Error() string {
	return fmt.Sprintf("Cannot process empty message")
}

type InvalidDataStorageManagerMessage struct {
	MalformedMessage []string
}

func (err *InvalidDataStorageManagerMessage) Error() string {
	return fmt.Sprintf("Cannot process malformed message: %s", err.MalformedMessage)
}

type MissingInputFile struct {
	InvalidFilename string
}

func (err *MissingInputFile) Error() string {
	return fmt.Sprintf("Unable to open missing file: %s", err.InvalidFilename)
}
