package controllers

import (
	auth "demiray.dev/goback/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	auth.Login(c)
}
