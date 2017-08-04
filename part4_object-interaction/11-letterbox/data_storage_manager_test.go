package storage

import "fmt"
import "testing"

func TestCanAllocateStruct(t *testing.T) {
	dataStorageManager := &DataStorageManager{}

	if dataStorageManager == nil {
		t.Error("dataStorageManager must not be nil!")
	}
}

func TestDataStorageManagerHasDispatchMethod(t *testing.T) {
	var iface interface{} = &DataStorageManager{}
	_, ok := iface.(interface {
		Dispatch(message []string) error
	})

	if !ok {
		t.Error("Struct DataStorageManagaer does not have any method 'Dispatch'!")
	}
}

func TestDispatchThrowsErrorOnEmptyMessage(t *testing.T) {
	dataStorageManager := &DataStorageManager{}
	message := []string{}
	err := dataStorageManager.Dispatch(message)

	if err == nil {
		t.Error("Dispatch method should have returned an error!")
	}

	var actualErr interface{} = err
	_, ok := actualErr.(*EmptyDataStorageManagerMessage)

	if !ok {
		t.Error(fmt.Sprintf("Dispatch method should have thrown EmptyDataStorageManagerMessage error, got %T", actualErr))
	}
}

func TestDispatchThrowsErrorOnMalformedInitMessage(t *testing.T) {
	dataStorageManager := &DataStorageManager{}
	initMessage := []string{"init"}
	err := dataStorageManager.Dispatch(initMessage)

	if err == nil {
		t.Error("Dispatch method should have returned an error!")
	}

	var expectedErr interface{} = err
	_, ok := expectedErr.(*InvalidDataStorageManagerMessage)

	if !ok {
		t.Error("Dispatch method did not throw expected error type!")
	}
}

func TestDispatchThrowsErrorOnMissingInputFile(t *testing.T) {
	dataStorageManager := &DataStorageManager{}

	initMessage := []string{"init", "foobar.txt"}
	err := dataStorageManager.Dispatch(initMessage)

	if err == nil {
		t.Error("Dispatch method should have thrown error!")
	}

	var expectedErr interface{} = err
	_, ok := expectedErr.(*MissingInputFile)

	if !ok {
		t.Error("Dispatch method should have thrown MissingInputFile error!")
	}

}

func TestFileContentIsSetOnExistingInputFile(t *testing.T) {
	dataStorageManager := &DataStorageManager{}

	initMessage := []string{"init", "test.txt"}
	err := dataStorageManager.Dispatch(initMessage)

	if err != nil {
		t.Error(fmt.Sprintf("Dispatch method should not have thrown an error, but got: %s", err.Error()))
	}

	if dataStorageManager.fileContent == "" {
		t.Error("File content must not be empty!")
	}
}

func TestFileContentReadMatchesExpectedContent(t *testing.T) {
	dataStorageManager := &DataStorageManager{}

	initMessage := []string{"init", "test.txt"}
	dataStorageManager.Dispatch(initMessage)

	if !matchesExpectedContent(dataStorageManager.fileContent) {
		t.Error(fmt.Sprintf("File content does not match expected value, got: %s", dataStorageManager.fileContent))
	}
}

func matchesExpectedContent(fileContent string) bool {
	return fileContent == "first line second line"
}
