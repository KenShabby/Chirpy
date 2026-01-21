package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8081"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	svr := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
