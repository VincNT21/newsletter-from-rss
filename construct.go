package main

import (
	"context"
	"fmt"
	"time"
)

type Items struct {
	Title       string
	Link        string
	Description string
	PubDate     time.Time
}

func getNewsletterData(gc GlobalConfig) (Newsletter, error) {
	// Create a Newletter struct
	newsletterData := Newsletter{
		MainTitle: gc.MainTitle,
		IntroText: gc.IntroText,
		OutroText: gc.OutroText,
		OutroLink: gc.OutroLink,
	}

	// Createan empty NewsletterCategory slice
	categoriesData := []NewsletterCategory{}

	// Iterate over given categories
	for _, cat := range gc.Categories {
		categoryData := NewsletterCategory{
			Title: cat.Title,
		}
		// If no sub categories
		if len(cat.SubCategories) == 0 {
			categoryData.HasSub = false
			// Check if link provided (not empty category)
			if cat.RssLink != "" {
				items, err := fetchAndSortRssData(cat)
				if err != nil {
					return newsletterData, err
				}
				categoryData.Items = items
			}

		} else {
			// If subcategories
			categoryData.HasSub = true
			subCategories := []NewsletterCategory{}
			for _, subCat := range cat.SubCategories {
				// Check if link provided (not empty category)
				subItems := []NewsletterItems{}
				if subCat.RssLink != "" {
					fetchItems, err := fetchAndSortRssData(subCat)
					if err != nil {
						return newsletterData, err
					}
					subItems = fetchItems
				}

				subCategories = append(subCategories, NewsletterCategory{
					Title:  subCat.Title,
					HasSub: false,
					Items:  subItems,
				})
			}
			categoryData.SubCategories = subCategories
		}

		categoriesData = append(categoriesData, categoryData)
	}

	// Set the categories data to Newsletter struct
	newsletterData.Categories = categoriesData

	return newsletterData, nil

}

func fetchAndSortRssData(cat Category) ([]NewsletterItems, error) {
	rssLink := cat.RssLink
	// Get the feed items
	rssfeed, err := fetchFeed(context.Background(), rssLink)
	if err != nil {
		return []NewsletterItems{}, fmt.Errorf("couldn't get rssfeed from %v: %w", rssLink, err)
	}
	rssItems := getItemsFromFeed(*rssfeed)

	// Filter items
	if cat.KeyWord != "" {
		rssItems = filterRssItemsByKeyword(rssItems, cat.KeyWord)
	}

	if cat.MaxItemsNumber != 0 {
		rssItems = filterRssItemsByMaxNumbers(rssItems, cat.MaxItemsNumber)
	}

	if cat.DaysInterval != 0 {
		rssItems = filterRssItemsByInterval(rssItems, cat.DaysInterval)
	}

	// Sort items by date
	rssItems = sortRssItemsByDate(rssItems)

	// Convert items to NewsletterItems and return them
	return convertItems(rssItems), nil
}
