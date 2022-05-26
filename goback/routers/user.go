package routers

import (
	"demiray.dev/goback/controllers"

	"github.com/gin-gonic/gin"
)

func UserAPIRoute(r *gin.RouterGroup) {
	r.GET("/login", controllers.Login)
}
