package firebaseUtil

import (
	"context"
	f "firebase.google.com/go"
	"firebase.google.com/go/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var authClient *auth.Client

func InitFirebaseService()  {
	opt := option.WithCredentialsFile("config/firebase.json")
	app, err := f.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	authClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}