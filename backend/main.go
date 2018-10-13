package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8000"},
	})
	app.Use(c)

	app.Get("/health", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "ok",
		})
	})

	app.Run(iris.Addr(":8080"))
}
