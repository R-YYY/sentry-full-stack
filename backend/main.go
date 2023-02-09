package main

import (
	"sentry-full-stack/backend/config"
	"sentry-full-stack/backend/route"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := route.InitRouter()
	config.InitConfig()

	gin.SetMode(viper.GetString("server.run_mode"))
	r.Run(viper.GetString("server.addr"))
}
