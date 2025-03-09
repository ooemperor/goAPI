package src

import (
	"github.com/gin-gonic/gin"
	"goAPI/src/api"
)

func InitRouter() *gin.Engine {
	ApiRouter := gin.Default()

	ApiRouter.GET("/user", api.UserApiObject.GetAll)
	ApiRouter.GET("user/:id", api.UserApiObject.GET)

	ApiRouter.DELETE("user/", api.UserApiObject.DeleteAll)
	ApiRouter.DELETE("user/:id", api.UserApiObject.DELETE)

	return ApiRouter
}
