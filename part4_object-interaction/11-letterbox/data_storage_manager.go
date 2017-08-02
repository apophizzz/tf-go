package storage

import (
	"fmt"
)

type DataStorageManager struct {
}

func (dsm *DataStorageManager) initialize(filePath string) {
}

func (dsm *DataStorageManager) Dispatch(message []string) error {
	if len(message) >= 1 {
		if message[0] == "init" {
			if len(message) != 2 {
				return &InvalidDataStorageManagerInitMessage{message}
			} else {
				return nil
			}
		} else {
			return nil
		}
	} else {
		return &EmptyDataStorageManagerMessage{}
	}
}

type EmptyDataStorageManagerMessage struct {
}

func (err *EmptyDataStorageManagerMessage) Error() string {
	return fmt.Sprintf("Cannot process empty message")
}

type InvalidDataStorageManagerInitMessage struct {
	MalformedInitMessage []string
}

func (err *InvalidDataStorageManagerInitMessage) Error() string {
	return fmt.Sprintf("Cannot process malformed init message: %s", err.MalformedInitMessage)
}
