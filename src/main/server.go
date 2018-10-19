package main

import (
	plg "../plugins/windows"
	"github.com/iris-contrib/middleware/cors"
	// "github.com/rs/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default() //router
	app.Use(cors.Default())

	app.Post("/softwareList", softwareListHndlr)

	app.Run(iris.Addr(":9000"))
}

func softwareListHndlr(ctx iris.Context) {
	swList := plg.SoftwareList()
	// ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(&swList)
}
