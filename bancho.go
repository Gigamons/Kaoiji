package main

import (
	"net/http"

	"cyanidee/bancho-go/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handle)
	http.ListenAndServe(":80", nil)
}
