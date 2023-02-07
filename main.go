package main

import (
	"log"
	"net/http"
	"rainbow/framework"
)

func main() {
	c := framework.NewCore()
	registerRouter(c)
	server := &http.Server{
		Handler: c,
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
