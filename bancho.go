package main

import (
	"net/http"

	"github.com/cyanidee/bancho-go/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handle)
	http.ListenAndServe(":80", nil)
}
