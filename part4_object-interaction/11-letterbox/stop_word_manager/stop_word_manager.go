package stopWords

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type StopWordManager struct {
	stopWords []string
}

func (swm *StopWordManager) Dispatch(message []string) ([]string, error) {
	if message[0] == "init" {
		swm.readStopWordsFromFile()
		return nil, nil
	} else if message[0] == "is_stop_word" {
		return nil, nil
	}
	return nil, &InvalidStopWordManagerMessage{message}
}

func (swm *StopWordManager) readStopWordsFromFile() {
	fileBytes, err := ioutil.ReadFile("stop_words.txt")
	exitOnError(err)
	swm.stopWords = strings.Split(strings.TrimSpace(string(fileBytes)), ",")
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to open file: stop_words.txt", err)
		os.Exit(1)
	}
}

type InvalidStopWordManagerMessage struct {
	invalidMessage []string
}

func (err *InvalidStopWordManagerMessage) Error() string {
	return fmt.Sprintf("Unable to understand message: %s", err.invalidMessage)
}
