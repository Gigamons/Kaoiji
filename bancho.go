package main

import (
	"fmt"
	"github.com/cyanidee/bancho-go/helpers"
	"log"
	"net/http"

	"github.com/cyanidee/bancho-go/handlers"
)

func main() {
	err, conf := helpers.ReadConfig()
	fmt.Println(err, conf)

	http.HandleFunc("/", handlers.Handle)
	log.Fatal(http.ListenAndServe(":80", nil))
}
