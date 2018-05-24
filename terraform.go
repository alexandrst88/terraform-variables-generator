package main

import (
	"regexp"
	"strings"
)

type terraformVars struct {
	Variables []string
}

func (t *terraformVars) matchVarPref(row, varPrefix string) {
	if strings.Contains(row, varPrefix) {
		pattern := regexp.MustCompile(`var.([a-z?0-9?_][a-z?0-9?_?-]*)`)
		match := pattern.FindAllStringSubmatch(row, 1)
		if len(match) != 0 {
			res := replacer.Replace(match[0][0])
			if !containsElement(t.Variables, res) {
				t.Variables = append(t.Variables, res)
			}
		}
	}
}
