package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func sendRequest() {
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
	for i := 0; i < 5; i++ {
		sendRequest()
	}

	fmt.Println("Sequential time taken: ", time.Since(start))
}
