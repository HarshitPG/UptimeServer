package Handler

import (
	"fmt"
	"net/http"
	"time"
)

var (
    lastPingStatus string
)

func init() {
	url := "https://zmt3q4-8080.csb.app/health" 

	
	pingURL := func() {
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
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running. Last ping status: %s\n", lastPingStatus)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}