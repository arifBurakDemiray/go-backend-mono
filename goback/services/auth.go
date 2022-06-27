package services

import (
	"errors"
	"net/http"
	"os"
	"time"

	libs "demiray.dev/goback/helpers"
	"demiray.dev/goback/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	defer libs.RecoverError(c)
	var (
		status             = 200
		msg                string
		responseData       = gin.H{}
		token              string
		userName, password string
		userModel          models.User
		err                error
		db                 *gorm.DB
	)

	db = libs.Connect()

	userName, _ = libs.GetQueryParam("user_name", c)
	password, _ = libs.GetQueryParam("password", c)
	resultFind := db.Where("email = ? AND password = ?", userName, password).First(&userModel)
	if resultFind.Error == nil || errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
		if resultFind.RowsAffected <= 0 {
			status = http.StatusUnauthorized
			msg = "incorrect user_name/pwd"
		} else {
			atClaims := jwt.MapClaims{}
			atClaims["user_id"] = userModel.ID
			atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
			at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			token, err = at.SignedString([]byte(os.Getenv("JWT_TOKEN")))
			if err != nil {
				status = 500
				msg = err.Error()
			}
		}
	} else {
		status = 500
	}
	if status == 200 {
		msg = "Success"
		responseData = gin.H{
			"status": status,
			"data":   token,
			"msg":    msg,
		}
	} else {
		if msg == "" {
			msg = "Error"
		}
		responseData = gin.H{
			"status": status,
			"msg":    msg,
		}
	}
	libs.APIResponseData(c, status, responseData)
}
