package api

import (
	"log"
	"time"

	"github.com/Mshivam2409/hls-streamer/internal/db"
	"github.com/dgrijalva/jwt-go"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/spf13/viper"
)

type Claims struct {
	UserID     string `json:"uid"`
	QuestionID string `json:"qid"`
	RequestID  string `json:"rid"`
	jwt.StandardClaims
}

func GenerateToken(rid string) (string, error) {
	token, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	dur, err := time.ParseDuration(viper.GetString("cache.expiry"))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	expirationTime := dur * time.Second
	err = db.GoStreamer.BadgerClient.Save(token, rid, expirationTime)
	if err != nil {
		return "", err
	}
	return token, err
}
