package main

import (
	"bufio"
	"html/template"
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

var replacer *strings.Replacer
var varPrefix = "var."
var tfFileExt = "*.tf"

var dstFile = "./variables.tf"
var varTemplate = template.Must(template.New("var_file").Parse(`{{range .}}
variable "{{ . }}" {
  description = ""
}
{{end}}`))

func init() {
	replacer = strings.NewReplacer(":", ".",
		"]", "",
		"}", "",
		"{", "",
		"\"", "",
		")", "",
		"(", "",
		"[", "",
		",", "",
		"var.", "",
		" ", "",
	)
}
func main() {
	if fileExists(dstFile) {
		userPromt()
	}

	tfFiles, err := getAllFiles(tfFileExt)
	if len(tfFiles) == 0 {
		log.Warn("No terraform files to proceed, exiting")
		os.Exit(0)
	}
	checkError(err)
	var wg sync.WaitGroup
	messages := make(chan string)
	wg.Add(len(tfFiles))
	t := &terraformVars{}

	for _, file := range tfFiles {
		go func(file string) {
			defer wg.Done()
			fileHandle, _ := os.Open(file)
			defer fileHandle.Close()
			fileScanner := bufio.NewScanner(fileHandle)
			for fileScanner.Scan() {
				messages <- fileScanner.Text()
			}
		}(file)
	}
	go func() {
		for text := range messages {
			t.matchVarPref(text, varPrefix)
		}
	}()
	wg.Wait()
	f, err := os.Create(dstFile)
	checkError(err)

	t.sortVars()
	err = varTemplate.Execute(f, t.Variables)
	checkError(err)
	log.Infof("Variables are generated to %q file", dstFile)

}
