package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:3000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
