package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/0xedb/go_client/client"
)

func main() {
	url := "http://localhost:2021/"
	fmt.Println("hello")

	var wg sync.WaitGroup

	time.AfterFunc(2*time.Second, func() {

		client.MakeRequest(url)
	})

	wg.Add(1)
	go client.StartServer(&wg)

	wg.Wait()
}
