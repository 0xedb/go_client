package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func getClient() *http.Client {
	client := &http.Client{}

	client.Timeout = 6 * time.Second

	return client
}

func MakeRequest(url string) {
	var client = getClient()
 
	res, err := client.Get(url)
	fmt.Println("making request...")

	if err != nil {
		log.Fatal("unable to make request ", err)
	}

	answer := bytes.Buffer{}
	answer.ReadFrom(res.Body)
	fmt.Println("got ", answer.String())
}

func getServer() *http.Server {
	return &http.Server{Addr: ":2021"}
}

func StartServer(wg *sync.WaitGroup) {
	defer wg.Done()
	server := getServer()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		res.Header().Add("Content-Type", "application/json")
		res.Write([]byte("{\"msg\":\"BYE!!!\"}"))
	})

	fmt.Println("starting server on", server.Addr)
	server.ListenAndServe()
}
