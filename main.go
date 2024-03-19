package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BoruTamena/server-go/server/server2"
)

func main() {

	http.HandleFunc("/receive", server2.Receive_data)

	fmt.Println("starting server...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
