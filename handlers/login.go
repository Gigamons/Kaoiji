package handlers

import (
	"github.com/Gigamons/Kaoiji/packets"
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
	
	ctx.Write(pw.GetBytes())
}