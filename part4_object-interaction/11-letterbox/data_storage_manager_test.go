package storage

import "fmt"
import "testing"

func TestCanAllocateStruct(t *testing.T) {
	dataStorageManager := &DataStorageManager{}

	if dataStorageManager == nil {
		t.Error("dataStorageManager must not be nil!")
	}
}

func TestDataStorageManagerHasInitMethod(t *testing.T) {
	dataStorageManager := &DataStorageManager{}
	var i interface{} = dataStorageManager
	_, ok := i.(interface {
		initialize(filePath string)
	})

	if !ok {
		t.Error("Struct DataStorageManager does not have any method 'init'!")
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
	_, ok := expectedErr.(*InvalidDataStorageManagerInitMessage)

	if !ok {
		t.Error("Dispatch method did not throw expected error type!")
	}
}
