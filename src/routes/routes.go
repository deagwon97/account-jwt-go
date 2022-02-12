package routes

import (
	"github.com/gin-gonic/gin"

	docs "account-jwt/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	accountRest "account-jwt/account/rest"
)

func Run(address string) error {

	router := gin.Default()

	docs.SwaggerInfo_swagger.BasePath = "/"

	v1 := router.Group("/")
	accountRest.AddAccountRoutes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router.Run(address)
}
