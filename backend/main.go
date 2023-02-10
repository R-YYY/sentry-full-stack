package main

import (
	"fmt"
	"sentry-full-stack/backend/config"
	"sentry-full-stack/backend/route"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	initSentry()

	r := route.InitRouter()
	r.Use(sentrygin.New(sentrygin.Options{}))
	config.InitConfig()
	gin.SetMode(viper.GetString("server.run_mode"))
	r.Run(viper.GetString("server.addr"))
}

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://adf8e01004144f7b8075d0d7ba7596bc@o4504648733032448.ingest.sentry.io/4504654755528704",
		Debug:            true,
		AttachStacktrace: true,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		fmt.Println("ERROR~~~~~~~~~~~~~~~~~")
	}
}
