package main

import (
	"log"
	"os"

	"demiray.dev/goback/middlewares"
	"demiray.dev/goback/models"
	"demiray.dev/goback/routers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	return ":8888"
}

var db *gorm.DB
var err error

func main() {
	dsn := "root:1qaz2WSX3edc>@tcp(127.0.0.1:3306)/goback?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Role{})
	port := getPort()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(location.Default())
	r.Use(middlewares.CORSMiddleware())
	rg := r.Group("goback")
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
