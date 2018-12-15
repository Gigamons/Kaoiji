package handlers

import (
	"cyanidee/bancho-go/helpers"
	"fmt"
	"net/http"
	"time"
)

func handleGet(response http.ResponseWriter, request *http.Request) {
	fmt.Println(helpers.GetArrayBytes([]int32{1, 2, 3, 4, 5}))
	fmt.Fprintf(response, "GET: %s", "hi")
}

func handlePost(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "POST: %s", time.Now())
}

// Handle handles the request fuck off VSCode
func Handle(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handleGet(response, request)

	case http.MethodPost:
		handlePost(response, request)

	}
}
