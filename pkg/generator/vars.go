package generator

import (
	"bufio"
	"html/template"
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/alexandrst88/terraform-variables-generator/pkg/utils"
)

var replacer *strings.Replacer
var varPrefix = "var."
var localPrefix = "local."

var varTemplate = template.Must(template.New("var_file").Funcs(template.FuncMap{"sub": sub}).Parse(`{{- $length := len .Variables -}}
{{- range $i, $v := .Variables -}}
{{ if $.VariablesDescription }}variable "{{ $v }}" {
  description = ""
}{{ else }}variable "{{ $v }}" {}{{ end }}
{{- if lt $i (sub $length 1) }}{{ "\n\n" }}{{ end -}}
{{ end -}}{{printf "\n"}}`))

func sub(a, b int) int { return a - b }

var localsTemplate = template.Must(template.New("locals_file").Parse(`locals { {{ range . }}
  {{ . }} ={{ end }}
}
`))

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
		"local.", "",
		" ", "",
	)
}

// Generate will write inputs to file
func Generate(tfFiles []string, varsDstFile string, localsDstFile string, varsDescription bool) {
	var wg sync.WaitGroup
	messages := make(chan string)
	wg.Add(len(tfFiles))
	t := &terraformVars{}
	t.VariablesDescription = varsDescription

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
			t.matchLocalPref(text, localPrefix)
		}
	}()
	wg.Wait()

	if len(t.Variables) > 0 {
		f, err := os.Create(varsDstFile)
		utils.CheckError(err)
		log.Infof("Variables are generated to %q file", varsDstFile)

		t.sort(t.Variables)
		err = varTemplate.Execute(f, t)
		utils.CheckError(err)
	}

	if len(t.Locals) > 0 {
		t.sort(t.Locals)
		localsFile, err := os.Create(localsDstFile)
		utils.CheckError(err)
		err = localsTemplate.Execute(localsFile, t.Locals)
		utils.CheckError(err)
		log.Infof("Locals are generated to %q file", localsDstFile)
	}
}
