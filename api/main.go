package main

import (
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	viper.SetConfigFile("config.json")
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("Error reading config file: %s", err.Error())
	//}
	_ = os.Setenv("SECRET", viper.GetString("jwt_secret"))
}

func main() {
	db.DB()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	port := os.Getenv("PORT")
	conn := "0.0.0.0:" + port


	log.Printf("Server running on %s", conn)
	log.Fatal(r.Run(conn))
}
