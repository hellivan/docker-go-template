package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Reads a YAML document from the values_in stream, uses it as values
// for the tpl_files templates and writes the executed templates to
// the out stream.
func ExecuteTemplates(values_in string, out io.Writer, tpl_files ...string) error {
	tpl, err := template.ParseFiles(tpl_files...)
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err)
	}

	dat, err := ioutil.ReadFile(values_in)
	if err != nil {
		return fmt.Errorf("Failed to read variables file: %v", err)
	}

	var values map[string]interface{}
	err = yaml.Unmarshal(dat, &values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err)
	}

	err = tpl.Execute(out, values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err)
	}
	return nil
}

func main() {
	err := ExecuteTemplates(os.Args[1], os.Stdout, os.Args[2:]...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}