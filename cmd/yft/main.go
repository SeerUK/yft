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
	tpl, err := template.ParseFiles(templates...)
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

	tpl.Funcs(template.FuncMap{
		"indent": tmplhelp.Indent,
	})

	err = tpl.Execute(output, values)
	if err != nil {
		return fmt.Errorf("error: Failed to parse template: %v", err)
	}

	return nil
}
