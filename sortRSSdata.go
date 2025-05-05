package main

import (
	"sort"
	"strings"
	"time"
)

func filterRssItemsByMaxNumbers(rssItems []Items, maxNumber int) []Items {
	if len(rssItems) <= maxNumber {
		return rssItems
	}
	return rssItems[:maxNumber]
}

func filterRssItemsByKeyword(rssItems []Items, keyword string) []Items {
	filteredItems := []Items{}

	// Iterate over items to filter if link contains keyword
	for _, item := range rssItems {
		if strings.Contains(item.Link, keyword) {
			filteredItems = append(filteredItems, item)
		}
	}

	return filteredItems
}

func filterRssItemsByInterval(rssItems []Items, daysInteval int) []Items {
	filteredItems := []Items{}

	// Set time comparators
	today := time.Now()

	// Check pub date
	for _, item := range rssItems {
		if item.PubDate.Before(today.AddDate(0, 0, -daysInteval)) {
			filteredItems = append(filteredItems, item)
		}
	}

	return filteredItems
}

func sortRssItemsByDate(rssItems []Items) []Items {

	// Sort by publication date
	sort.Slice(rssItems, func(i, j int) bool {
		// Check if either time is the zero value
		if rssItems[i].PubDate.IsZero() && !rssItems[j].PubDate.IsZero() {
			// Put item with zero time at the end
			return false
		}
		if !rssItems[i].PubDate.IsZero() && rssItems[j].PubDate.IsZero() {
			// Put item with zero time at the end
			return true
		}
		if rssItems[i].PubDate.IsZero() && rssItems[j].PubDate.IsZero() {
			// If both are zero, sort by title
			return rssItems[i].Title < rssItems[j].Title
		}

		// Normal case: sort by pub date (most recent first)
		return rssItems[i].PubDate.After(rssItems[j].PubDate)
	})

	return rssItems
}
