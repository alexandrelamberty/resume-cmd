/*
Resume-cmd generate resumes.

Usage:

	resume-cmd [flags] [path ...]

The flags are:

	-i
	    The resume YAML file.

	-t
	    The resume template file.

	-o
	    The output filename.

When gofmt reads from standard input, it accepts either a full Go program
or a program fragment. A program fragment must be a syntactically
valid declaration list, statement list, or expression. When formatting
such a fragment, gofmt preserves leading indentation as well as leading
and trailing spaces, so that individual sections of a Go program can be
formatted by piping them through gofmt.
*/
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func toBase64(file string) string {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding
}

var (
	i_flag string
	t_flag string
	o_flag string
	i_dir  string
	t_dir  string
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func main() {
	// Command line argument
	flag.StringVar(&i_flag, "i", "resume.yml", "YAML resume file")
	flag.StringVar(&t_flag, "t", "tpl.gohtml", "GoHTML template file or directory")
	flag.StringVar(&o_flag, "o", "resume.html", "Output HTML file")
	flag.Parse()

	// Check empty arguments
	if !isFlagPassed("i") {
		fmt.Println("-i flag not passed")
	}
	if !isFlagPassed("t") {
		fmt.Println("-t flag not passed")
		os.Exit(1)
	}
	// Check if t_flag is a dir
	if info, err := os.Stat(t_flag); err == nil && info.IsDir() {
		fmt.Println("-t flag must be a file")
		os.Exit(1)
	}

	if _, err := os.Stat(i_flag); err == nil && os.IsNotExist(err) {
		fmt.Println("-i flag must be a file")
		os.Exit(1)
	}

	// Retrieve the resume file path
	i_dir := filepath.Dir(i_flag)
	t_dir := filepath.Dir(t_flag)

	// Read the Yaml data file
	data := map[string]interface{}{}
	buf, err := ioutil.ReadFile(i_flag)
	check(err)
	err = yaml.Unmarshal(buf, &data)
	check(err)

	// Convert image file to base64 string
	// Create the file path relative to the -i flag
	fp := filepath.Join(i_dir, data["Picture"].(string))
	// Convert to base64
	img := toBase64(fp)
	// Replace picture filename with base64 string
	data["Picture"] = img

	// Create a file on the system
	// FIXME: read o flag or default
	if o_flag == "" {
		o_flag = "resume.html"
	}
	out, err := os.Create(o_flag)
	check(err)
	defer out.Close()

	// Templates
	// FIXME Could be a single file or directory
	type any = map[string]interface{}
	// fmt.Println(t_flag)
	t := template.Must(template.New(filepath.Base(t_flag)).Funcs(template.FuncMap{
		// Convert unsafe string to template.URL #ZgotmplZ
		"Picture": func(u any) template.URL {
			img := u["Picture"].(string)
			return template.URL(img)
		},
		// FIXME: dir || file
	}).ParseGlob(t_dir + "/*"))
	err = t.Execute(out, data)
	check(err)

	fmt.Println(ColorBlue, "Resume converted!", string(ColorReset))
}
