package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	lastPingStatus string
)

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

func main() {
	url := "https://zmt3q4-8080.csb.app/health"
	startPing(url)

	e := echo.New()
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Server is running. Last ping status: %s\n", lastPingStatus))
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Server is running. Last ping status: %s\n", lastPingStatus))
	})
	e.ServeHTTP(w, r)
}
