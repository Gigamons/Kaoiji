package handlers

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Kaoiji/helpers"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
	"github.com/valyala/fasthttp"
)

func login_request(ctx *fasthttp.RequestCtx) {
	presence := objects.NewPresence()
	ctx.Response.Header.Set("cho-token", presence.Token)

	pw := packets.PacketWriter{}
	pw.ProtocolNegotiation(19)

	body := ctx.PostBody()
	if body == nil {
		login_exception(ctx)
		return
	}

	loginRequest := helpers.ParseLoginRequest(body)
	if loginRequest == nil {
		login_outdated(ctx)
		return
	}

	pw.LoginReply(-1)
	pw.Announce("Hello Golang")

	ctx.Write(pw.GetBytes())
}

func login_exception(ctx *fasthttp.RequestCtx) {
	pw := packets.PacketWriter{}
	pw.LoginReply(consts.LoginException)
	ctx.Write(pw.GetBytes())
}

func login_outdated(ctx *fasthttp.RequestCtx) {
	pw := packets.PacketWriter{}
	pw.LoginReply(consts.LoginClientOutdated)
	ctx.Write(pw.GetBytes())
}