package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func generateNewsletter(data Newsletter, dateFormat string) (string, error) {

	// Get template path
	templatePath := getFilePath("html_template.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("couldn't parse files for template: %w", err)
	}

	// Get proper today date for stamping the file name
	today := time.Now()
	todayParsed := today.Format(dateFormat)

	// Create output file
	fileName := "newsletter_" + todayParsed + ".html"
	outputPath := getFilePath(fileName)
	f, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("couldn't create file: %w", err)
	}
	defer f.Close()

	// Execute template with feed items pre-organized
	err = tmpl.Execute(f, data)
	return fileName, err
}
