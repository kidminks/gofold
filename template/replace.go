package template

import "strings"

type ReplaceMap struct {
	Module   string
	ValueMap map[string]string
}

func (r *ReplaceMap) CreateValueMap() {
	r.ValueMap = make(map[string]string)
	if r.Module != "" {
		r.ValueMap["{module}"] = r.Module
	}
}

func (r *ReplaceMap) Replace(input string) string {
	output := input
	for k, v := range r.ValueMap {
		output = strings.ReplaceAll(output, k, v)
	}
	return output
}
