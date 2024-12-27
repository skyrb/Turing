package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	maxRetries     = 3
	retryInterval = 500 * time.Millisecond
	timeout       = 2 * time.Second
)

func getData(url string, wg *sync.WaitGroup) ([]byte, error) {
	defer wg.Done()

	for i := 0; i < maxRetries; i++ {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		client := &http.Client{Timeout: timeout}
		resp, err := client.Do(req)
		if err != nil {
			if resp != nil {
				resp.Body.Close()
			}
			fmt.Printf("Retry %d failed for %s: %v\n", i+1, url, err)
			time.Sleep(retryInterval)
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Retry %d failed for %s: Status code %d\n", i+1, url, resp.StatusCode)
			time.Sleep(retryInterval)
			continue
		}

		return nil, nil
	}

	return nil, errors.New("maximum retries exceeded")
}

func main() {
	urls := []string{"https://example.com", "https://example.org", "https://example.net"}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go getData(url, &wg)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}