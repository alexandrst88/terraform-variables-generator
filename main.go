package main

import (
	"bufio"
	"html/template"
	"os"
	"regexp"
	"strings"
)

var replacer *strings.Replacer

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

var varTemplate = template.Must(template.New("var_file").Parse(`{{ range . }} variable "{{ . }}" {
   description  = ""
}
{{end}}
`))

var var_prefix = "var."

type TerraformVars struct {
	Variables []string
}

func contains(slice []string, value string) bool {
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

func main() {
	var parsed_vars []string
	fileHandle, _ := os.Open("test.tf")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), var_prefix) {
			pattern := regexp.MustCompile(`var.([a-z?_]+)`)
			match := pattern.FindAllStringSubmatch(fileScanner.Text(), 1)
			if len(match) != 0 {
				res := replacer.Replace(match[0][0])
				if !contains(parsed_vars, res) {
					parsed_vars = append(parsed_vars, res)
				}
			}
		}
	}
	err := varTemplate.Execute(os.Stdout, parsed_vars)
	if err != nil {
		panic(err)
	}
}
