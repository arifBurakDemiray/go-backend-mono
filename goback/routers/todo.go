package routers

import (
	"demiray.dev/goback/controllers"

	"github.com/gin-gonic/gin"
)

func TodoAPIRoute(r *gin.RouterGroup) {
	r.GET("/tasks", controllers.GetAllTask)
	r.POST("/tasks", controllers.CreateTask)
}
