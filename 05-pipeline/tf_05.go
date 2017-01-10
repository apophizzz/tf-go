package main

import (
	"io/ioutil"
	"regexp"
	"strings"
	"log"
)

func readFile(filePath string) string {
	fileData, err := ioutil.ReadFile(filePath)

	if (err != nil) {
		panic(err)
	}

	return string(fileData)
}

func filterCharsAndNormalize(data string) string {
	regexp, err := regexp.Compile("[\\W_]+")

	if (err != nil) {
		panic(err)
	}

	return strings.ToLower(regexp.ReplaceAllLiteralString(data, " "))
}

func main() {
	log.Print(filterCharsAndNormalize(readFile("input.txt")))
}
