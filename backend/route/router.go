package route

import (
	"sentry-full-stack/backend/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", api.GetIndex)

	router.GET("/sentry", api.GetSentry)

	return router
}
