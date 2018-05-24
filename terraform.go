package main

import (
	"regexp"
	"sort"
	"strings"
)

type terraformVars struct {
	Variables []string
}

func (t *terraformVars) matchVarPref(row, varPrefix string) {
	if strings.Contains(row, varPrefix) {
		pattern := regexp.MustCompile(`var.([a-z?A-Z?0-9?_][a-z?A-Z?0-9?_?-]*)`)
		match := pattern.FindAllStringSubmatch(row, -1)
		for _, m := range match {
			res := replacer.Replace(m[0])
			if !containsElement(t.Variables, res) {
				t.Variables = append(t.Variables, res)
			}
		}
	}
}

func (t *terraformVars) sortVars() {
	sort.Sort(sort.StringSlice(t.Variables))
}
