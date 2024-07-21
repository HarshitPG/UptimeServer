package handler

import (
	"fmt"
	"net/http"
	"time"
)

var lastPingStatus string

func pingURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		lastPingStatus = fmt.Sprintf("Error pinging URL: %s", err)
		fmt.Printf("Error pinging URL: %s\n", err)
		return
	}
	defer resp.Body.Close()
	lastPingStatus = fmt.Sprintf("Pinged %s - Status Code: %d", url, resp.StatusCode)
	fmt.Printf("Pinged %s - Status Code: %d\n", url, resp.StatusCode)
}

func startPing(url string) {
	pingURL(url)
	ticker := time.NewTicker(3 * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				pingURL(url)
			}
		}
	}()
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/status" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Server is running. Last ping status: %s\n", lastPingStatus)
}

func Handler() {
	url := "https://zmt3q4-8080.csb.app/health"
	startPing(url)

	http.HandleFunc("/status", statusHandler)
	fmt.Println("Server is starting...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
