package parsing

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//EBNF represents a classification of a subset of a program text
type EBNF struct {
	tag      string //EBNF classification
	content  string //the string that is classified
	children []EBNF //child classifications
}

type grammar struct {
	tag     string
	options []grammar
}

//creates a string representation of the grammar file given
func parse(grammarFile string) grammar {
	b, err := ioutil.ReadFile(grammarFile)
	if err != nil {
		return grammar{tag: "error", options: nil}
	}
	body := fmt.Sprintf("%s", b)
	lines := strings.Split(body, ";")
	var gmap map[string]*grammar
	for lines, line := decouple(lines); len(lines) > 0; lines, line = decouple(lines) {
		tokens := GTok(line)
		if tokens[1] != ":=" { //assert that this is a valid line
			return grammar{tag: "error", options: nil}
		}
		var g grammar
		gmap[tokens[0]] = &g
	}
	return grammar{tag: "Program", options: nil}
}

func decouple(lines []string) ([]string, string) {
	o := lines[len(lines)-1]
	return lines[:len(lines)-1], o
}
