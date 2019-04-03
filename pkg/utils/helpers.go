package utils

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// CheckError is a convenient wrapper for error check
func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ContainsElement check if slice of strings contains provided string
func ContainsElement(slice []string, value string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, s := range slice {
		if value == s {
			return true
		}
	}
	return false
}

// UserPromt will ask user if file needs to be overridden
func UserPromt(dstFile string) {
	var response string
	log.Warnf("File %q already exists, type 'yes' if you want replace it", dstFile)
	fmt.Print("-> ")
	_, err := fmt.Scanln(&response)
	CheckError(err)
	if response != "yes" {
		os.Exit(0)
	}
}
