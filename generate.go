package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

type Newsletter struct {
	MainTitle  string
	IntroText  string
	OutroText  string
	OutroLink  string
	Categories []NewsletterCategory
}

type NewsletterCategory struct {
	Title         string
	HasSub        bool
	SubCategories []NewsletterCategory
	Items         []NewsletterItems
}

type NewsletterItems struct {
	Title       string
	Link        string
	Description string
	PubDate     string
}

func generateNewsletter(data Newsletter) error {

	// Get template path
	templatePath := getFilePath("html_template.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("couldn't parse files for template: %w", err)
	}

	// Get proper today date for stamping the file name
	today := time.Now()
	todayParsed := fmt.Sprintf("%02d-%02d-%d", today.Year(), today.Month(), today.Day())

	// Create output file
	fileName := "newsletter_" + todayParsed + ".html"
	outputPath := getFilePath(fileName)
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("couldn't create file: %w", err)
	}
	defer f.Close()

	// Execute template with feed items pre-organized
	return tmpl.Execute(f, data)
}
