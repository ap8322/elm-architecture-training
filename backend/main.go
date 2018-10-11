package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/health", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "ok",
		})
	})

	app.Run(iris.Addr(":8080"))
}
