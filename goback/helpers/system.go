package libs

import (
	"log"
	"os"
	"strings"

	"demiray.dev/goback/models"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// RecoverError func
func RecoverError(c *gin.Context) {
	if r := recover(); r != nil {
		responseData := gin.H{
			"status": 500,
			"msg":    r,
		}
		c.JSON(500, responseData)
		return
	}
}

//TODO seperate db related funcs to db helper

func Connect() *gorm.DB {
	var db *gorm.DB
	var err error
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	return db
}

func InitDbTables() {

	db := Connect()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Role{})

}

// APIResponseData func
func APIResponseData(c *gin.Context, status int, responseData gin.H) {
	responseType := c.Request.Header.Get("ResponseType")
	if responseType == "application/xml" {
		c.XML(status, responseData)
	} else {
		c.JSON(status, responseData)
	}
}

func GetUserIDFromToken(token string) string {
	var userID string
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})
	if err == nil {
		if t.Valid {
			sUserID, ok := claims["user_id"].(string)
			if ok {
				userID = sUserID
			}
		}
	}
	return userID
}

// GetQueryParam func
func GetQueryParam(param string, c *gin.Context) (string, bool) {
	vParam, sParam := c.GetQuery(param)
	vParamLower, sParamLower := c.GetQuery(strings.ToLower(param))
	if sParam {
		return vParam, sParam
	}
	if sParamLower {
		return vParamLower, sParamLower
	}
	return "", false
}
