package generator

import (
	"bufio"
	"os"
	"testing"

	"github.com/alexandrst88/terraform-variables-generator/pkg/utils"
)

var testExtFile = "*.mock"

func TestContainsElement(t *testing.T) {
	testSlice := []string{"Terraform", "Puppet", "Ansible"}
	if utils.ContainsElement(testSlice, "Chef") {
		t.Error("Should return false, but return true")
	}
}

func TestGetAllFiles(t *testing.T) {
	files, err := utils.GetAllFiles(testExtFile)
	utils.CheckError(err)
	if len(files) == 0 {
		t.Error("Should found at least one file")
	}
}

func TestMatchVariable(t *testing.T) {
	ter := &terraformVars{}
	var messages []string

	file, err := utils.GetAllFiles(testExtFile)
	utils.CheckError(err)

	fileHandle, _ := os.Open(file[0])
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		messages = append(messages, fileScanner.Text())
	}
	for _, text := range messages {
		ter.matchVarPref(text, varPrefix)
	}
	if len(ter.Variables) != 5 {
		t.Errorf("Should return five variable. but returned %d", len(ter.Variables))
		t.Errorf("Variables found: %s", ter.Variables)
	}

}
