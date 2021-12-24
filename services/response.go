package services

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type resTemplate struct {
	Status		string			`json:"status"`
	Message		string			`json:"message"`
	Data		interface{}		`json:"data"`
}

func SendJSONResponse(ctx *fasthttp.RequestCtx, data interface{}, status, message string)  {
	ctx.SetContentType("application/json")

	if data == nil {
		data = map[int]int{}
	}
	if status == "" {
		status = "Success"
	}
	if message == "" {
		message = "Success"
	}

	res := &resTemplate{
		Status: status,
		Message: message,
		Data: data,
	}
	out, e := json.Marshal(res)
	if e != nil {
		panic(e)
	}

	ctx.SetBody(out)
}