package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RanchoCooper/distributed-object-storage.git/chapter01/objects"
)

func main() {
	http.HandleFunc(objects.DefaultPrefix, objects.Handler)
	err := http.ListenAndServe(os.Getenv(objects.ListenAddress), nil)
	if err != nil {
		log.Fatal(err)
	}
}
