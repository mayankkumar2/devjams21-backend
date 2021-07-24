package sentryUtil

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"os"
)

func InitSentry()  {
	sentryUrl := os.Getenv("SENTRY_URL")
	fmt.Println("[INFO] initializing sentry:", sentryUrl)
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryUrl,
	})
	if err != nil {
		logrus.Fatalf("sentry.Init: %s", err)
	}
}
