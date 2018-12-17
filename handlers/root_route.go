package handlers

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func handleGet(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "GET: %s", time.Now())
}

func handlePost(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("cho-protocol", "19")
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.Response.Header.Set("Keep-Alive", "timeout=5, max=100")
	ctx.Response.Header.Set("Content-Type", "application/octet-stream; charset=UTF-8")

	if string(ctx.UserAgent()) != "osu!" {
		handleGet(ctx) // it's not osu! we wont handle that. (as for now)
		return
	}

	if string(ctx.Request.Header.Peek("osu-token")) == "" {
		login_request(ctx)
		return
	}

}

// HandleRoot handles the request fuck off VSCode
func HandleRoot(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "GET":
		handleGet(ctx)

	case "POST":
		handlePost(ctx)
	}
}
