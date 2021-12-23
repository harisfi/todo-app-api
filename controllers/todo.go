package controllers

import "github.com/valyala/fasthttp"

func GetAllTodo(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	a := ctx.FormValue("activity_group_id")
	ctx.SetBody([]byte("GetAllTodo: you're in " + string(ctx.RequestURI()) +
						", activity_group_id = " + string(a)))
}

func GetOneTodo(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("GetOneTodo: you're in " + string(ctx.RequestURI())))
}

func CreateTodo(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("CreateTodo: you're in " + string(ctx.RequestURI())))
}

func DeleteTodo(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("DeleteTodo: you're in " + string(ctx.RequestURI())))
}

func UpdateTodo(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("UpdateTodo: you're in " + string(ctx.RequestURI())))
}
