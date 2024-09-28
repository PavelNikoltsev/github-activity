package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func ExtractPages(linkHeader string) (*int, *int, error) {
	links := strings.Split(linkHeader, ",")
	var currentPage, lastPage int

	for _, link := range links {
		// Trim whitespace and extract the URL and rel value
		link = strings.TrimSpace(link)
		parts := strings.Split(link, ";")
		if len(parts) < 2 {
			continue // Skip if there's no rel
		}

		// Extract the URL
		urlStr := strings.Trim(parts[0], "<>")
		u, err := url.Parse(urlStr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse URL: %s", urlStr)
		}

		// Get query parameters
		query := u.Query()
		pageStr := query.Get("page")
		if pageStr != "" {
			// Parse the page number
			pageNum, err := strconv.Atoi(pageStr)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid page number: %s", pageStr)
			}

			// Determine if this is the current or last page based on rel
			if strings.Contains(parts[1], `rel="next"`) {
				currentPage = pageNum - 1
			} else if strings.Contains(parts[1], `rel="last"`) {
				lastPage = pageNum
			}
		}
	}
	return &currentPage, &lastPage, nil
}
