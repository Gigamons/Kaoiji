package handlers

import (
	"fmt"
	"github.com/cyanidee/bancho-go/packets"
	"net/http"
)

func login_request(response http.ResponseWriter, request *http.Request) {
	/*
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {

	}
	*/

	response.Header().Set("cho-protocol", "19")
	response.Header().Set("Connection", "keep-alive")
	response.Header().Set("Keep-Alive", "timeout=5, max=100")
	response.Header().Set("Content-Type", "application/octet-stream; charset=UTF-8")
	response.Header().Add("cho-token", "test")

	pw := packets.PacketWriter{}
	pw.LoginReply(-1)
	pw.Announce("Hello Golang")

	if err := pw.WriteBytes(response); err != nil {
		fmt.Println(err)
	}
}