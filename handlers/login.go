package handlers

import (
	"fmt"
	"github.com/cyanidee/bancho-go/packets"
	"github.com/valyala/fasthttp"
)

func login_request(ctx *fasthttp.RequestCtx) {
	/*
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {

	}
	*/

	ctx.Response.Header.Set("cho-token", "test")

	pw := packets.PacketWriter{}
	pw.LoginReply(-1)
	pw.Announce("Hello Golang")

	if err := pw.WriteBytes(ctx); err != nil {
		fmt.Println(err)
	}
}