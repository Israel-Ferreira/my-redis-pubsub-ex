package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/my-redis-pubsub/pkg/handlers"
)

func main() {
	fmt.Println("Redis Pub/Sub")

	port := ":4500"

	log.Println("Server started on Port ", port)

	http.HandleFunc("/cars", handlers.HandleCreateCarRequest)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln("Error on Start Server: ", err.Error())
	}

}
