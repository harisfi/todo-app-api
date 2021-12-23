package controllers

import "github.com/valyala/fasthttp"

func GetAllActivity(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("GetAllActivity: you're in " + string(ctx.RequestURI())))
}

func GetOneActivity(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("GetOneActivity: you're in " + string(ctx.RequestURI())))
}

func CreateActivity(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("CreateActivity: you're in " + string(ctx.RequestURI())))
}

func DeleteActivity(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("DeleteActivity: you're in " + string(ctx.RequestURI())))
}

func UpdateActivity(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("UpdateActivity: you're in " + string(ctx.RequestURI())))
}
