package fio

import (
	"bytes"
	"os"
	"testing"
)

func TestNewFileIOManager(t *testing.T) {
	fileName := "testfile.txt"

	// Clean up the file if it already exists
	defer os.Remove(fileName)

	fio, err := NewFileIOManager(fileName)
	if err != nil {
		t.Fatalf("Error creating FileIO: %v", err)
	}
	defer fio.Close()

	if fio.fd == nil {
		t.Errorf("Expected a valid file descriptor, but got nil")
	}
}

func TestFileIO_ReadWrite(t *testing.T) {
	fileName := "testfile.txt"

	// Clean up the file if it already exists
	defer os.Remove(fileName)

	fio, err := NewFileIOManager(fileName)
	if err != nil {
		t.Fatalf("Error creating FileIO: %v", err)
	}
	defer fio.Close()

	expectedData := []byte("Hello, World!")

	// Write data to the file
	n, err := fio.Write(expectedData)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}
	if n != len(expectedData) {
		t.Errorf("Expected to write %d bytes, but wrote %d bytes", len(expectedData), n)
	}

	// Read data from the file
	readData := make([]byte, len(expectedData))
	n, err = fio.Read(readData, 0)
	if err != nil {
		t.Fatalf("Error reading from file: %v", err)
	}
	if n != len(expectedData) {
		t.Errorf("Expected to read %d bytes, but read %d bytes", len(expectedData), n)
	}
	if !bytes.Equal(readData, expectedData) {
		t.Errorf("Read data doesn't match expected data")
	}
}

func TestFileIO_Sync(t *testing.T) {
	fileName := "testfile.txt"

	// Clean up the file if it already exists
	defer os.Remove(fileName)

	fio, err := NewFileIOManager(fileName)
	if err != nil {
		t.Fatalf("Error creating FileIO: %v", err)
	}
	defer fio.Close()

	// Perform a sync operation
	err = fio.Sync()
	if err != nil {
		t.Errorf("Error syncing file: %v", err)
	}
}
