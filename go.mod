// +heroku install ./api/main.go
// +heroku goVersion go1.16
module github.com/GDGVIT/devjams21-backend

go 1.16

require (
	cloud.google.com/go/storage v1.18.2 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/appleboy/gin-jwt/v2 v2.7.0
	github.com/bketelsen/crypt v0.0.4 // indirect
	github.com/getsentry/sentry-go v0.11.0
	github.com/gin-contrib/size v0.0.0-20211002110825-4c208e0712f3
	github.com/gin-gonic/gin v1.7.4
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/moby/moby v20.10.10+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.9.0
	github.com/tidwall/gjson v1.10.2 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/net v0.0.0-20211020060615-d418f374d309 // indirect
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359 // indirect
	golang.org/x/tools v0.1.7 // indirect
	google.golang.org/api v0.59.0
	google.golang.org/genproto v0.0.0-20211026145609-4688e4c4e024 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	gorm.io/driver/postgres v1.2.0
	gorm.io/gorm v1.22.0
)
