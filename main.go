package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		// URL of the webpage to check
		url := "https://www.nike.com/in/launch/t/air-jordan-1-mid-se-white-black"

		// Fetch the webpage
		response, err := http.Get(url)
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

		// Check if the webpage contains the text "Sold Out"
		if strings.Contains(string(body), "Sold Out") {
			fmt.Println("The webpage contains 'Sold Out'")
			sendNotification()
		} else {
			fmt.Println("The webpage does not contain 'Sold Out'")
		}

		// Sleep for 1 hour before checking again
		time.Sleep(time.Hour)
	}
}

func sendNotification() {
	// Execute the shell script to send a notification
	cmd := exec.Command("sh", "-p", "osascript -e 'display notification \"Sold Out\" with title \"Nike Alert\"'")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
		return
	}	
}

