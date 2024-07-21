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

    go func() {
		for {
            select {
            case <-ticker.C:
                pingURL()
            }
        }
	}() 

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go server is running!")
	})

	
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is up and running! Pinging URL: %s", url)
	})


	fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
