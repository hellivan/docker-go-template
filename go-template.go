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
		return fmt.Errorf("Failed to parse variables file as YAML: %v", err)
	}

	err = tpl.Execute(out, values)
	if err != nil {
		return fmt.Errorf("Failed to process template files: %v", err)
	}
	return nil
}

func printUsage(returnCode int) {
     fmt.Printf("Usage: %v VAR_FILE [FILE]...\n", os.Args[0])
     fmt.Printf("\nVAR_FILE\tFile in YAML format which contains variables definitions for the template\n")
     fmt.Printf("FILE\t\tList of template files which should be processed using the specified variables\n\n")
     os.Exit(returnCode)
}

func main() {
     
     if len(os.Args) < 3 {
     	printUsage(1)
     }

     err := ExecuteTemplates(os.Args[1], os.Stdout, os.Args[2:]...)
     if err != nil {
     	log.Println(err)
	os.Exit(1)
     }
}