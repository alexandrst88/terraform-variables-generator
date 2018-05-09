package main

import (
	"os"
	"testing"
)

var mockFile = "mock.tf"

func TestContainsElement(t *testing.T) {
	testSlice := []string{"Terraform", "Puppet", "Ansible"}
	if containsElement(testSlice, "Chef") {
		t.Error("Should return false, but return true")
	}
}

func TestGetAllFiles(t *testing.T) {
	var file, err = os.Create(mockFile)
	checkError(err)
	defer file.Close()
	files, err := getAllFiles(tfFileExt)
	checkError(err)
	if len(files) == 0 {
		t.Error("Should found at least one file")
	}
	defer os.Remove(mockFile)
}
