package routers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func SetupRouter() func(ctx *fasthttp.RequestCtx) {

	req := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			ctx.SetContentType("application/json")
			m, _ := json.Marshal(map[string]string{"message": "Welcome to API TODO"})
			ctx.SetBody(m)
		default:
			ctx.Error("404 Not Found", fasthttp.StatusNotFound)
		}
	}

	return req
}