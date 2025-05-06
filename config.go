package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GlobalConfig struct {
	MainTitle      string     `json:"main_title"`       // The main title of your newsletter
	IntroText      string     `json:"intro_text"`       // The intro text below the title
	OutroText      string     `json:"outro_text"`       // The outro text at the end of the newsletter
	OutroLink      string     `json:"outro_link"`       // The outro link below the outro text
	OutroLinkTitle string     `json:"outro_link_title"` // The text for outro link
	DateFormat     string     `json:"date_format"`      // The wanted format for date parsing
	Categories     []Category `json:"categories"`       // A list of categories
}

type Category struct {
	Title          string     `json:"title"`         // The category title
	RssLink        string     `json:"rss_link"`      // The category RSS feed link. Leave empty for an empty category
	KeyWord        string     `json:"keyword"`       // The category key word, as it will appear in link, if needed. Leave empty otherwise
	DaysInterval   int        `json:"days_interval"` // The duration in days within the category. Leave empty if needed
	MaxItemsNumber int        `json:"max_items"`     // The maximum number of items within the category. Leave empty if needed
	SubCategories  []Category `json:"subcategories"` // A list of subcategories. Leave empty if needed
}

func getGlobalConfig() (GlobalConfig, error) {
	filepath := getFilePath("config.json")
	byteData, err := os.ReadFile(filepath)
	if err != nil {
		return GlobalConfig{}, fmt.Errorf("could not read %s file: %v", filepath, err)
	}
	var gc GlobalConfig
	err = json.Unmarshal(byteData, &gc)
	if err != nil {
		return GlobalConfig{}, fmt.Errorf("could not unmarshal data: %v", err)
	}
	return gc, nil
}
