package main

import (
	"errors"
	"os"
	"testing"
	"time"
)

// TODO: Break up teests into logging
// ... And actual functionality (main.go)

func TestGenerateLogFile(t *testing.T) {
	GenerateLogFile("log_test.json")

	if _, err := os.Stat("log_test.json"); err == nil {
	} else {
		t.Error("Error, unable to create logfile", err)
	}

	testCleanup()
}

func TestCreateExampleFiles(t *testing.T) {
	GenerateExampleFiles()

	if _, err := os.Stat("example.txt"); err == nil {
	} else {
		t.Error("Error, unable to create logfile", err)
	}

	testCleanup()
}

func TestCreateFile(t *testing.T) {
	filepath := "./test_path/file.json"

	CreateFile(filepath)

	if _, err := os.Stat("./test_path/file.json"); err == nil {
	} else {
		t.Error("Error, file not being created", err)
	}

	testCleanup()
}

func TestDeleteFile(t *testing.T) {
	filepath := "./test_path/file.json"

	// Create file so it can be deleted
	CreateFile(filepath)

	DeleteFile(filepath)

	_, err := os.Open(filepath)

	if errors.Is(err, os.ErrNotExist) {
	} else {
		t.Error("File was not successfully deleted")
	}

	testCleanup()
}

func TestLogNetworkRequest(t *testing.T) {
	fileName := "log_test.json"

	GenerateLogFile(fileName)

	data := NetworkRequestEvent{
		UserName:           "Rico",
		ProcessName:        "NetworkRequest",
		CommandLine:        "NetworkRequest",
		Protocol:           "HTTP",
		DestinationAddress: "server.com",
		DestinationPort:    "80",
		SourceAddress:      "localhost",
		SourcePort:         "8080",
		DataAmount:         10, // MB
		Timestamp:          time.Now(),
	}
	// TODO: you could check for the presence of the log in the JSON file..
	// This applies to other logging functions as well
	LogNetworkRequest(data, fileName)
	testCleanup()
}

// TODO: Should we check for presence of file before writing?
// Tests the user can write a new process
// to the test_log.json file
func TestLogProcessStart(t *testing.T) {
	fileName := "log_test.json"

	GenerateLogFile(fileName)

	data := ProcessStartEvent{
		UserName:    "Rico",
		ProcessName: "ProcessStarted",
		CommandLine: "--arg",
		Timestamp:   time.Now(),
	}

	LogProcessStart(data, fileName)

	// TODO: Check if file is there to complete the test
	testCleanup()
}

func TestProcessStart(t *testing.T) {
	args := []string{"--a", "--b", "--c"}
	// path := "./example_executables/example_executable"
	path := "example_executables/example_executable"
	ProcessStart(path, args)
	t.Error("Incomplete process start")
	// Process
}

func TestLoggingMultipleStartProcesses(t *testing.T) {
	fileName := "log_test.json"

	GenerateLogFile(fileName)

	data := ProcessStartEvent{
		UserName:    "Rico",
		ProcessName: "ProcessStarted",
		CommandLine: "--arg",
		Timestamp:   time.Now(),
	}

	LogProcessStart(data, fileName)

	data2 := ProcessStartEvent{
		UserName:    "Rico2",
		ProcessName: "ProcessStarted",
		CommandLine: "--arg",
		Timestamp:   time.Now(),
	}

	LogProcessStart(data2, fileName)

	// TODO: check if the files are there to complete the test
	// TODO: testCleanup(filename) is needed to ensure cleanups work
	// properly

	testCleanup()
}

// TODO: Add test completion/failed scenarios
func TestLogFileChange(t *testing.T) {
	fileName := "log_test.json"

	GenerateLogFile(fileName)

	data := FileChangeEvent{
		UserName:    "Rico2",
		FilePath:    "users/egrueter/exec",
		Descriptor:  "create",
		ProcessName: "FileCreated",
		CommandLine: "--create",
		ProcessId:   123,
		Timestamp:   time.Now(),
	}

	LogFileChange(data, fileName)

	data2 := FileChangeEvent{
		UserName:    "Rico3",
		FilePath:    "users/egrueter/exec",
		Descriptor:  "delete",
		ProcessName: "FileCreated",
		CommandLine: "--create",
		ProcessId:   123,
		Timestamp:   time.Now(),
	}

	LogFileChange(data2, fileName)
	testCleanup()
}

// Clean up log_test.json file
// after every test run
func testCleanup() {
	os.Remove("log_test.json")
	os.Remove("example.txt")
	os.RemoveAll("./test_path/")
}
