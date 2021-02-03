package main

import (
	"github.com/kataras/iris/v12"
)

type request struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type response struct {
	ID      uint64 `json:"id"`
	Message string `json:"message"`
}

func main() {
	app := iris.New()
	app.Handle("PUT", "/users/{id:uint64}", updateUser)
	app.Listen(":8888")
}

func updateUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint64("id")

	var req request
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	resp := response{
		ID:      id,
		Message: req.Firstname + " " + req.Lastname + " updated successfully!",
	}

	ctx.JSON(resp)
}
