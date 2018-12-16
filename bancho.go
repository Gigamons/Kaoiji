package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cyanidee/bancho-go/handlers"
	"github.com/cyanidee/bancho-go/helpers"
)

func main() {
	if _, err := helpers.ConnectMySQL("test", 54, "hitler", "password", "hitlersleftball"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(helpers.DBConn.Ping())

	http.HandleFunc("/", handlers.Handle)
	log.Fatal(http.ListenAndServe(":80", nil))
}
