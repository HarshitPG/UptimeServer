package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://zmt3q4-8080.csb.app/health" 

	
	pingURL := func() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error pinging URL: %s\n", err)
			return
		}
		defer resp.Body.Close()
		fmt.Printf("Pinged %s - Status Code: %d\n", url, resp.StatusCode)
	}


	pingURL()

	ticker := time.NewTicker(3 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			pingURL()
		}
	}
}
