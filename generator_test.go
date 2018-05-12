package main

import (
	"bufio"
	"os"
	"testing"
)

var mockFile = "tf_configuration.mock"
var testExtFile = "*.mock"

func TestContainsElement(t *testing.T) {
	testSlice := []string{"Terraform", "Puppet", "Ansible"}
	if containsElement(testSlice, "Chef") {
		t.Error("Should return false, but return true")
	}
}

func TestGetAllFiles(t *testing.T) {
	files, err := getAllFiles(testExtFile)
	checkError(err)
	if len(files) == 0 {
		t.Error("Should found at least one file")
	}
}

func TestMatchVariable(t *testing.T) {
	ter := &terraformVars{}
	var messages []string

	file, err := getAllFiles(testExtFile)
	checkError(err)

	fileHandle, _ := os.Open(file[0])
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		messages = append(messages, fileScanner.Text())
	}
	for _, text := range messages {
		ter.matchVarPref(text, varPrefix)
	}
	if len(ter.Variables) != 1 {
		t.Errorf("Should return one variable. but returned %d", len(ter.Variables))
	}

}
