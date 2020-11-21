package generator

import (
	"regexp"
	"sort"
	"strings"

	"github.com/alexandrst88/terraform-variables-generator/pkg/utils"
)

type terraformVars struct {
	Variables []string
	Locals    []string
}

func (t *terraformVars) matchVarPref(row, varPrefix string) {
	if strings.Contains(row, varPrefix) {
		pattern := regexp.MustCompile(`var.([a-z?A-Z?0-9?_][a-z?A-Z?0-9?_?-]*)`)
		match := pattern.FindAllStringSubmatch(row, -1)
		for _, m := range match {
			res := replacer.Replace(m[0])
			if !utils.ContainsElement(t.Variables, res) {
				t.Variables = append(t.Variables, res)
			}
		}
	}
}

func (t *terraformVars) matchLocalPref(row, localPrefix string) {
	if strings.Contains(row, localPrefix) {
		pattern := regexp.MustCompile(`local.([a-z?A-Z?0-9?_][a-z?A-Z?0-9?_?-]*)`)
		match := pattern.FindAllStringSubmatch(row, -1)
		for _, m := range match {
			res := replacer.Replace(m[0])
			if !utils.ContainsElement(t.Locals, res) {
				t.Locals = append(t.Locals, res)
			}
		}
	}
}

func (t *terraformVars) sort(vars []string) {
	sort.Strings(vars)
}
