package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	PubDateAlt  string `xml:"dc:date"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Init client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error with GET request: %w", err)
	}
	req.Header.Set("User-Agent", "github.com/VincNT21/newsletter-from-rss")

	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error with GET request: %w", err)
	}
	defer resp.Body.Close()

	// Handle resp -> data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error with ReadAll: %w", err)
	}

	// Decode data into xml
	result := RSSFeed{}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error with xml Unmarshal: %w", err)
	}

	// Decode escaped HTML entities
	result.Channel.Description = html.UnescapeString(result.Channel.Description)
	result.Channel.Title = html.UnescapeString(result.Channel.Title)
	for i, item := range result.Channel.Items {
		result.Channel.Items[i].Description = html.UnescapeString(item.Description)
		result.Channel.Items[i].Title = html.UnescapeString(item.Title)
	}

	return &result, nil
}

func getItemsFromFeed(feed RSSFeed) []Items {
	items := []Items{}
	// Iterate over feed items
	for _, item := range feed.Channel.Items {
		// Check for pubdate
		parsedPubDate := time.Time{}
		if item.PubDate != "" {
			parsedPubDate = parsePubDate(item.PubDate)
		} else if item.PubDateAlt != "" {
			parsedPubDate = parsePubDate(item.PubDateAlt)
		}
		items = append(items, Items{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			PubDate:     parsedPubDate,
		})
	}
	return items
}
