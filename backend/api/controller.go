package api

import (
	"github.com/Gocyber-world/gocyber-base/response"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	response.Ok(c)
}

func GetSentry(c *gin.Context) {
	response.OkWithData("sentry", c)
}
