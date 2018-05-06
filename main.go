package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var replacer *strings.Replacer
var tf_file_ext = ".tf"
var var_prefix = "var."
var varTemplate = template.Must(template.New("var_file").Parse(`{{ range . }} variable "{{ . }}" {
	description  = ""
 }
 {{end}}
 `))

type TerraformVars struct {
	Variables []string
}

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

func getAllFiles(ext string) []string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
				fmt.Printf("Found file: %q\n", f.Name())
			}
		}
		return nil
	})
	return files
}

func (t *TerraformVars) matchVarPref(row, var_prefix string) {
	if strings.Contains(row, var_prefix) {
		pattern := regexp.MustCompile(`var.([a-z?_]+)`)
		match := pattern.FindAllStringSubmatch(row, 1)
		if len(match) != 0 {
			res := replacer.Replace(match[0][0])
			if !containsElement(t.Variables, res) {
				t.Variables = append(t.Variables, res)
			}
		}
	}
}

func main() {
	tf_files := getAllFiles(tf_file_ext)
	var wg sync.WaitGroup
	messages := make(chan string)
	wg.Add(len(tf_files))
	t := &TerraformVars{}

	for _, file := range tf_files {
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
			t.matchVarPref(text, var_prefix)
		}
	}()
	wg.Wait()
	err := varTemplate.Execute(os.Stdout, t.Variables)
	if err != nil {
		panic(err)
	}
}
