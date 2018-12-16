package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func handleGet(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "GET: %s", time.Now())
}

func handlePost(response http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(response, "POST: %s", time.Now())

	fmt.Println(request.Header.Get("osu-token"))
	fmt.Println(request.UserAgent())
	if request.Header.Get("osu-token") == "" &&
		request.UserAgent() == "osu!"{
		login_request(response, request)
	}

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
