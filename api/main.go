package main

import (
	"github.com/GDGVIT/devjams21-backend/api/middleware"
	"github.com/GDGVIT/devjams21-backend/api/router"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/GDGVIT/devjams21-backend/pkg/firebaseUtil"
	"github.com/GDGVIT/devjams21-backend/pkg/sentryUtil"
	"github.com/getsentry/sentry-go"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %s\n", err.Error())
	} else {
		_ = os.Setenv("SECRET", viper.GetString("SECRET"))
		_ = os.Setenv("PORT", viper.GetString("PORT"))
		_ = os.Setenv("DATABASE_URL", viper.GetString("DATABASE_URL"))
		_ = os.Setenv("DEPLOYMENT", viper.GetString("DEPLOYMENT"))
		_ = os.Setenv("SENTRY_URL", viper.GetString("SENTRY_URL"))
		_ = os.Setenv("PAYLOAD_SIZE", viper.GetString("PAYLOAD_SIZE"))
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	db.DB()
	firebaseUtil.InitFirebaseService()
	sentryUtil.InitSentry()
	defer sentry.Flush(2 * time.Second)

	payloadSizeLimit, err := strconv.Atoi(os.Getenv("PAYLOAD_SIZE"))
	if err != nil {
		payloadSizeLimit = 10 // 10 MB
	}

	r := gin.Default()
	gin.New()
	r.Use(middleware.CORSMiddleware())

	r.Use(limits.RequestSizeLimiter(int64(payloadSizeLimit)))

	api := r.Group("api")
	if os.Getenv("DEPLOYMENT") == "PUBLIC" {
		router.RegisterPublicRoutes(api)
	} else if os.Getenv("DEPLOYMENT") == "ADMIN" {
		router.RegisterAdminRoutes(api)
		router.RegisterPublicRoutes(api)
	}

	port := os.Getenv("PORT")
	conn := "0.0.0.0:" + port
	log.Printf("Server running on %s", conn)
	log.Fatal(r.Run(conn))
}
