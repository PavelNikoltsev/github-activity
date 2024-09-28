package main

import (
	"encoding/json"
	"fmt"
	"github-activity/models"
	"github-activity/utils"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	run()
}

func run() {
	page := 1
	limit := 10
	var username string
	var input string
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	resp, err := request(username, page, limit)
	if err != nil {
		log.Fatalf("Error getting events: %v", err)
	}
	if err := models.PrintEvents(resp.Events); err != nil {
		log.Fatalf("Error printing events: %v", err)
	}
	for {
		if resp.Page != nil && resp.TotalPages != nil {
			fmt.Printf("Page %d of %d\n", *resp.Page, *resp.TotalPages)
			fmt.Print("Enter page number (or 'q' to quit): ")
			fmt.Scan(&input)
			if input == "q" {
				break
			}
			page, err = strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Invalid page number: %v\n", err)
				continue
			}
			if page < 1 || page > *resp.TotalPages {
				fmt.Printf("Invalid page number: %d\n", page)
				continue
			}
			resp, err = request(username, page, limit)
			if err != nil {
				fmt.Printf("Error getting events: %v\n", err)
				continue
			}
			if err := models.PrintEvents(resp.Events); err != nil {
				log.Fatalf("Error printing events: %v", err)
			}
		} else {
			fmt.Print("Press Enter to perform a new request or type 'q' to quit: ")
			fmt.Scan(&input)
			if input == "q" {
				break
			}
			resp, err = request(username, page, limit)
			if err != nil {
				fmt.Printf("Error getting events: %v\n", err)
				continue
			}
			if err := models.PrintEvents(resp.Events); err != nil {
				log.Fatalf("Error printing events: %v", err)
			}
		}
	}
}

type response struct {
	Page       *int
	TotalPages *int
	Events     []models.Event
}

func request(username string, page, limit int) (*response, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events?page=%d&per_page=%d", username, page, limit)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error performing GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}

	response := &response{}
	linkHeader := resp.Header.Get("link")
	if linkHeader != "" {
		response.Page, response.TotalPages, err = utils.ExtractPages(linkHeader)
		if err != nil {
			return nil, err
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response.Events)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return response, nil
}
