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

var varTemplate = template.Must(template.New("var_file").Parse(`{{range .}}
variable "{{ . }}" {
  description = ""
}
{{end}}`))

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
func Generate(tfFiles []string, varsDstFile string, localsDstFile string) {
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
			t.matchLocalPref(text, localPrefix)
		}
	}()
	wg.Wait()

	if len(t.Variables) > 0 {
		f, err := os.Create(varsDstFile)
		utils.CheckError(err)
		log.Infof("Variables are generated to %q file", varsDstFile)

		t.sort(t.Variables)
		err = varTemplate.Execute(f, t.Variables)
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
