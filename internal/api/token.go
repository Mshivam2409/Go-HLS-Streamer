package api

import (
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
	expirationTime := time.Duration(viper.GetInt("cache.expiry")) * time.Minute
	err = db.GoStreamer.BadgerClient.Save(token, rid, expirationTime)
	if err != nil {
		return "", err
	}
	return token, err
}
