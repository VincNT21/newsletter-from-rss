package main

import "time"

type Items struct {
	Title       string
	Link        string
	Description string
	PubDate     time.Time
}

type Newsletter struct {
	MainTitle      string
	IntroText      string
	OutroText      string
	OutroLink      string
	OutroLinkTitle string
	Categories     []NewsletterCategory
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
