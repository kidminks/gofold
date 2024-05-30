package template

import "strings"

type ReplaceMap struct {
	Module   string
	Input    string
	Output   string
	ValueMap map[string]string
}

func (r *ReplaceMap) CreateValueMap() {
	r.ValueMap = make(map[string]string)
	if r.Module != "" {
		r.ValueMap["{module}"] = r.Module
	}
}

func (r *ReplaceMap) Replace() {
	output := r.Input
	for k, v := range r.ValueMap {
		output = strings.ReplaceAll(output, k, v)
	}
	r.Output = output
}
