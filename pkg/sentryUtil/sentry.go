package sentryUtil

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"os"
)

func InitSentry()  {
	sentryUrl := os.Getenv("SENTRY_URL")
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryUrl,
	})
	if err != nil {
		logrus.Fatalf("sentry.Init: %s", err)
	}
}
