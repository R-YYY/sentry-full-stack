package api

import (
	"fmt"

	"github.com/Gocyber-world/gocyber-base/response"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	response.Ok(c)
}

func GetSentry(c *gin.Context) {
	// ctx := context.Background()
	// span := sentry.StartSpan(ctx, "doWork", sentry.TransactionName("testtttttttt"))
	// span.Finish()
	hub := sentry.GetHubFromContext(c.Request.Context())
	headers := c.Request.Header
	fmt.Println(headers)
	fmt.Println(hub)
	// hub.Scope().SetTag("Trace ID", )
	// hub.Scope().SetTransaction("/sentry")

	response.OkWithData("sentry", c)
}
