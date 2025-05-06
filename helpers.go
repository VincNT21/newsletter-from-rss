package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func getFilePath(filename string) string {
	// Check for environment variable
	if filename == "html_template.html" {
		htlmPath := os.Getenv("TEMPLATE_PATH")
		if htlmPath != "" {
			return filepath.Join(htlmPath, filename)
		}
	} else if strings.Contains(filename, "newsletter") {
		outputPath := os.Getenv("OUTPUT_PATH")
		if outputPath != "" {
			return filepath.Join(outputPath, filename)
		}
	} else if filename == "app.log" {
		logPath := os.Getenv("LOG_PATH")
		if logPath != "" {
			return filepath.Join(logPath, filename)
		}
	}

	outputDir := ""

	// Check if OS used is Windows or else
	if runtime.GOOS == "windows" {
		// For Windows builds, use executable directory
		execPath, err := os.Executable()
		if err == nil {
			outputDir = filepath.Dir(execPath)
		}
	} else {
		outputDir, _ = os.Getwd()
	}

	// If we couldn't determine a directory, fall back to current directory
	if outputDir == "" {
		outputDir = "."
	}

	return filepath.Join(outputDir, filename)
}

func parsePubDate(pubDateString string) time.Time {
	// Some possible formats for pubDate
	formats := []string{
		time.RFC1123Z,
		time.RFC3339,
		time.RFC1123,
		"2006-01-02T15:04:05-07:00",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"02 Jan 2006 15:04:05 MST",
		"Mon, 2 Jan 2006 15:04:05 -0700",
		// Add more formats if needed
	}

	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, pubDateString)
		if err == nil {
			// Successfully parsed
			return parsedTime
		}
		// Else, continue with next format
	}

	// If none of the formats worked, return an empty/Zero value
	return time.Time{}
}

func convertItems(items []Items, dateFormat string) []NewsletterItems {
	newsletterItems := []NewsletterItems{}

	for _, item := range items {
		newsletterItems = append(newsletterItems, NewsletterItems{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			PubDate:     "Published on " + item.PubDate.Format(dateFormat),
		})
	}

	return newsletterItems
}
