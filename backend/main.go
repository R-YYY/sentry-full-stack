package main

import (
	"log"
	"sentry-full-stack/backend/config"
	"sentry-full-stack/backend/route"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	initSentry()

	r := route.InitRouter()
	config.InitConfig()

	gin.SetMode(viper.GetString("server.run_mode"))
	r.Run(viper.GetString("server.addr"))
}

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "https://989ada487cb947edab59b82c218b31fa@o4504648733032448.ingest.sentry.io/4504648734867456",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
		AttachStacktrace: true,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			// Here you can inspect/modify non-transaction events (for example, errors) before they are sent.
			// Returning nil drops the event.
			log.Printf("BeforeSend event [%s]", event.EventID)
			return event
		},
		BeforeSendTransaction: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			// Here you can inspect/modify transaction events before they are sent.
			// Returning nil drops the event.
			if strings.Contains(event.Message, "test-transaction") {
				// Drop the transaction
				return nil
			}
			event.Message += " [example]"
			log.Printf("BeforeSendTransaction event [%s]", event.EventID)
			return event
		},
		// Enable tracing
		EnableTracing: true,
		// Specify either a TracesSampleRate...
		TracesSampleRate: 1.0,
		// ... or a TracesSampler
		TracesSampler: sentry.TracesSampler(func(ctx sentry.SamplingContext) float64 {
			// As an example, this custom sampler does not send some
			// transactions to Sentry based on their name.
			hub := sentry.GetHubFromContext(ctx.Span.Context())
			name := hub.Scope().Transaction()
			if name == "GET /favicon.ico" {
				return 0.0
			}
			if strings.HasPrefix(name, "HEAD") {
				return 0.0
			}
			// As an example, sample some transactions with a uniform rate.
			if strings.HasPrefix(name, "POST") {
				return 0.2
			}
			// Sample all other transactions for testing. On
			// production, use TracesSampleRate with a rate adequate
			// for your traffic, or use the SamplingContext to
			// customize sampling per-transaction.
			return 1.0
		}),
	})
	if err != nil {
		return
	}
}
