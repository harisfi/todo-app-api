package services

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func ExtractRequests(ctx *fasthttp.RequestCtx) map[string]interface{} {
	var req map[string]interface{}

	if len(ctx.PostBody()) > 0 {
		if json.Valid(ctx.PostBody()) {
			if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
				panic(err)
			}
		}
	}
	return req
}