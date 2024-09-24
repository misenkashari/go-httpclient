package main

import (
	"fmt"
	"go-httpclient/core"
)

func main() {
	client := core.NewClient("https://httpbin.org")

	respChannel := make(chan string)

	for i := 0; i < 10; i++ {
		go func(c chan string, i int) {
			resp, err := client.Get("/get")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			result := fmt.Sprintf("Call %d GET Response Status: %d\n", i, resp.StatusCode)
			body := fmt.Sprintf("Call %d GET Response Body: %s\n", i, resp.Body)
			c <- result
			c <- body
		}(respChannel, i)

		fmt.Println(<-respChannel)
	}
}
