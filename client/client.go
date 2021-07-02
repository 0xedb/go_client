package client

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getClient() *http.Client {
	client := &http.Client{}

	// client.
	client.Timeout = 10 * time.Second

	return client
}

func MakeRequest(url string) {
	var client = getClient()

	// client.Do(
	res, err := client.Get(url)
	fmt.Println("making request...")

	fmt.Println("TIMEOUT", client.Timeout)

	if err != nil {
		log.Fatal("unable to make request ", err)
	}

	answer := []byte{}
	res.Body.Read(answer)
	fmt.Println(string(answer))
}

func getServer() *http.Server {
	return &http.Server{Addr: ":2021"}
}

func StartServer() {
	server := getServer()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(10 * time.Second)
		res.Header().Add("Content-Type", "application/json")
		res.Write([]byte("{\"msg\":\"goodbye\"}"))
	})

	fmt.Println("starting server on", server.Addr)
	server.ListenAndServe()
}
