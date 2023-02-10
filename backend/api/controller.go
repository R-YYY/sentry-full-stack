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
	span := StartSpanFromGinContext(c, "sentry test")
	for i := 0; i < 5; i++ {
		childOp := span.StartChild(fmt.Sprintf("slow op: %v", i))
		time.Sleep(100 * time.Millisecond)
		childOp.Finish()
	}
	// traceId, _ := hex.DecodeString(c.Request.Header.Get("traceId"))
	// span.TraceID = sentry.TraceID(traceId)
	sentry.ContinueFromHeaders(c.Request.Header.Get("sentry-trace"), "")
	span.Finish()
	response.OkWithData(span.TraceID, c)
	fmt.Println(span.TraceID)
	sentry.Flush(time.Second)
}

func StartSpanFromGinContext(c *gin.Context, op string) *sentry.Span {
	span, ok := c.Value("_sentry_gin_span").(*sentry.Span)
	if ok && span != nil {
		return span.StartChild(op, sentry.ContinueFromHeaders(c.Request.Header.Get("sentry-trace"), ""))
	}
	return sentry.StartSpan(c, op, sentry.ContinueFromHeaders(c.Request.Header.Get("sentry-trace"), ""))
}
