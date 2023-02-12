package api

import (
	"fmt"
	"time"

	"github.com/Gocyber-world/gocyber-base/response"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	response.Ok(c)
}

func GetSentry(c *gin.Context) {
	span := StartRootSpan(c, "http.server", "GET /api/sentry")
	for i := 0; i < 5; i++ {
		childOp := span.StartChild(fmt.Sprintf("slow op: %v", i))
		time.Sleep(100 * time.Millisecond)
		childOp.Finish()
	}
	span.Finish()
	response.OkWithData(span.TraceID, c)
	fmt.Println(span.TraceID)
	sentry.Flush(time.Second)
}

func StartRootSpan(c *gin.Context, op string, name string) *sentry.Span {
	span := sentry.StartTransaction(c, name, sentry.ContinueFromTrace(c.Request.Header.Get("sentry-trace")))
	span.Op = op
	return span
}
