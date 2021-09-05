// +heroku install ./api/main.go
// +heroku goVersion go1.16
module github.com/GDGVIT/devjams21-backend

go 1.16

require (
	cloud.google.com/go v0.88.0 // indirect
	cloud.google.com/go/firestore v1.5.0 // indirect
	cloud.google.com/go/storage v1.16.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/appleboy/gin-jwt/v2 v2.6.5-0.20210827121450-79689222c755
	github.com/getsentry/sentry-go v0.11.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/validator/v10 v10.7.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.3.0
	github.com/jackc/pgx/v4 v4.12.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	google.golang.org/api v0.51.0
	google.golang.org/genproto v0.0.0-20210722135532-667f2b7c528f // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.12
)
