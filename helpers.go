package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func containsElement(slice []string, value string) bool {
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

func userPromt() {
	var response string
	log.Warnf("File %q already exists, type yes if you want replace", dstFile)
	fmt.Print("-> ")
	_, err := fmt.Scanln(&response)
	checkError(err)
	if response != "yes" {
		os.Exit(0)
	}
}
