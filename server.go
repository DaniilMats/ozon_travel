package main

import (
	"log"
	"net/http"

	"github.com/DaniilMats/ozon_travel/endpointHandlers"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", endpointHandlers.FormHandler)

	log.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
