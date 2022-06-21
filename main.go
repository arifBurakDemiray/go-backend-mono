package main

import (
	"os"

	libs "demiray.dev/goback/helpers"
	"demiray.dev/goback/middlewares"
	"demiray.dev/goback/routers"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func getPort() string {
	p := os.Getenv("HOST_PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + os.Getenv("PORT")
}

func main() {

	libs.InitDbTables()

	port := getPort()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(location.Default())
	r.Use(middlewares.CORSMiddleware())
	rg := r.Group("api")
	rg.Use(middlewares.CORSMiddleware())
	{
		routers.UserAPIRoute(rg)
	}
	rg.Use(middlewares.CORSMiddleware(), middlewares.ValidateToken())
	{
		routers.TodoAPIRoute(rg)
	}
	r.Run(port)
}
