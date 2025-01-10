package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Error      error
}

func main() {
	// CLI flags
	baseURL := flag.String("url", "", "Base URL to fuzz (e.g. http://example.com/FUZZ)")
	wordlist := flag.String("w", "", "Path to wordlist file")
	workers := flag.Int("t", 10, "Number of concurrent workers")
	rateLimit := flag.Int("rate", 10, "Requests per second")
	flag.Parse()

	if *baseURL == "" || *wordlist == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Read wordlist
	words, err := loadWordlist(*wordlist)
	if err != nil {
		fmt.Printf("Error loading wordlist: %v\n", err)
		os.Exit(1)
	}

	// Create channels
	jobs := make(chan string, *workers)
	results := make(chan Result, *workers)
	rateLimiter := time.NewTicker(time.Second / time.Duration(*rateLimit))
	defer rateLimiter.Stop()

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go worker(jobs, results, rateLimiter.C, &wg)
	}

	// Send jobs
	go func() {
		for _, word := range words {
			url := strings.Replace(*baseURL, "FUZZ", word, -1)
			jobs <- url
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error: %s - %v\n", result.URL, result.Error)
			continue
		}
		if result.StatusCode != 404 {
			fmt.Printf("Found: %s - Status: %d\n", result.URL, result.StatusCode)
		}
	}
}

func worker(jobs <-chan string, results chan<- Result, rateLimiter <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for url := range jobs {
		<-rateLimiter // Rate limiting
		resp, err := client.Get(url)
		if err != nil {
			results <- Result{URL: url, Error: err}
			continue
		}
		results <- Result{URL: url, StatusCode: resp.StatusCode}
		resp.Body.Close()
	}
}

func loadWordlist(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}
