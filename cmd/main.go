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

func toBase6(file string) string {
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
	flag.StringVar(&i_flag, "i", "resume.yml", "resume file")
	flag.StringVar(&t_flag, "t", "/resume-theme/*", "template directory")
	flag.StringVar(&o_flag, "o", "html", "output file")
	flag.Parse()
	// Check empty arguments
	if !isFlagPassed("i") {
		fmt.Println("-i flag not passed")
	}
	// Create a file on the system
	out, err := os.Create("resume.html")
	defer out.Close()
	check(err)

	// Read the Yaml data file
	data := map[string]interface{}{}
	buf, err := ioutil.ReadFile(i_flag)
	check(err)
	err = yaml.Unmarshal(buf, &data)
	check(err)

	// Convert image file to base64 string
	// Read picture filename
	img := toBase6(data["Picture"].(string))
	// Replace picture filename with base64 string
	data["Picture"] = img

	// Templateing
	// t, err := template.ParseGlob("templates/blue/*")
	check(err)
	type any = map[string]interface{}
	t := template.Must(template.New("_.html").Funcs(template.FuncMap{
		// Convert unsafe string to template.URL #ZgotmplZ
		"Picture": func(u any) template.URL {
			img := u["Picture"].(string)
			return template.URL(img)
		},
	}).ParseGlob(t_flag+"/*")) // TODO Uniq file or directory

	err = t.Execute(out, data)
	check(err)
	fmt.Println(ColorBlue, "Resume converted!", string(ColorReset))
}
