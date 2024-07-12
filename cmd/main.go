package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"jordanSoldOutChecker/internal/notify"
)

type Config struct {
	URL           string `json:"url"`
	CheckInterval int    `json:"checkInterval"`
	SearchText    string `json:"searchText"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig("config/controlCenter.json")
	if err != nil {
		fmt.Printf("Error loading config file: %v\n", err)
		return
	}

	for {
		// Fetch the webpage
		response, err := http.Get(config.URL)
		if err != nil {
			fmt.Printf("Error fetching URL: %v\n", err)
			return
		}
		defer response.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			return
		}

		// Check if the webpage contains the specified text
		if strings.Contains(string(body), config.SearchText) {
			fmt.Printf("The webpage contains '%s'\n", config.SearchText)
			notify.SendNotification(config.SearchText)
		} else {
			fmt.Printf("The webpage does not contain '%s'\n", config.SearchText)
		}

		// Sleep for the specified interval before checking again
		time.Sleep(time.Duration(config.CheckInterval) * time.Second)
	}
}
