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
        return
    }
    defer resp.Body.Close()
    lastPingStatus = fmt.Sprintf("Pinged %s - Status Code: %d", url, resp.StatusCode)
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

func PingHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    url := "https://zmt3q4-8080.csb.app/health"
    startPing(url)
    fmt.Fprintf(w, "Server is running. Last ping status: %s\n", lastPingStatus)
}

func main() {
    url := "https://zmt3q4-8080.csb.app/health"
    startPing(url)

    http.HandleFunc("/", PingHandler)
    http.ListenAndServe(":3000", nil)
}
