package handlers

import (
	"fmt"
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Kaoiji/helpers"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
	"github.com/Gigamons/Shared/sutilities"
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

	userId, err := sutilities.GetUserId(loginRequest.UserName)
	if err != nil {
		login_exception(ctx)
		fmt.Println(err)
		return
	}
	if userId <= 0 {
		login_failed(ctx, "This Username doesn't exists!")
		return
	}

	user, err := sutilities.GetUser(userId)
	if err != nil {
		login_exception(ctx)
		fmt.Println(err)
		return
	}
	if user == nil {
		pw.LoginReply(consts.LoginException)
		pw.Announce("This User doesn't exists! (Exception, ask a Developer for help)")
		ctx.Write(pw.GetBytes())
		return
	}

	if !user.IsPassword(loginRequest.PassMD5) {
		login_failed(ctx, "This Password is incorrect!")
		return
	}

	presence.User = user
	presence.UserStatus, err = sutilities.GetUserStatus(user.Id)
	if err != nil {
		login_exception(ctx)
		fmt.Println(err)
		return
	}
	presence.UTCOffset = loginRequest.UTCOffset


	/* TODO: Check for ban, get country, setup leaderboard */
	pw.LoginReply(consts.LoginReply(user.Id))
	pw.Presence(presence.GetUserPresence())


	objects.AppendPresence(presence)
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

func login_failed(ctx *fasthttp.RequestCtx, msg string) {
	pw := packets.PacketWriter{}
	pw.LoginReply(consts.LoginClientOutdated)
	pw.Announce(msg)
	ctx.Write(pw.GetBytes())
}
