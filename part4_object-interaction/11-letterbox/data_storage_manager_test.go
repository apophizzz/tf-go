package storage

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
		dispatch(message []string)
	})

	if !ok {
		t.Error("Struct DataStorageManagaer does not have any method 'dispatch'!")
	}
}
