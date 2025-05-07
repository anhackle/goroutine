package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func sendRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest("GET", "http://testasp.vulnweb.com/", nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status Code: ", resp.StatusCode)

}

func main() {
	// Target: Send multiple requests to a server
	// sequentially vs simultaneously
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sendRequest(&wg)
	}
	wg.Wait()
	fmt.Println("Sequential time taken: ", time.Since(start))
}
