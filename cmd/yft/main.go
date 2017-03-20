package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/SeerUK/yft/pkg/tmplhelp"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	err := render(os.Stdin, os.Stdout, os.Args[1:]...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// render the given template files using the given input, parsed as YAML for variables.
func render(input io.Reader, output io.Writer, templates ...string) error {
	funcMap := template.FuncMap{
		"indent": tmplhelp.Indent,
	}

	tpl, err := template.New("").Funcs(funcMap).ParseFiles(templates...)
	if err != nil {
		return fmt.Errorf("error: Failed to parse template(s): %v", err)
	}

	raw, err := ioutil.ReadAll(input)
	if err != nil {
		return fmt.Errorf("error: Failed reading variables from standard input: %v", err)
	}

	var values map[string]interface{}
	err = yaml.Unmarshal(raw, &values)
	if err != nil {
		return fmt.Errorf("error: Failed unmarshalling YAML input: %v", err)
	}

	for _, t := range tpl.Templates() {
		err = tpl.ExecuteTemplate(output, t.Name(), values)
		if err != nil {
			return fmt.Errorf("error: Failed to parse template '%s': %v", t.Name(), err)
		}
	}

	return nil
}
