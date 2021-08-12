package api

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID     string `json:"uid"`
	QuestionID string `json:"qid"`
	RequestID  string `json:"rid"`
	jwt.StandardClaims
}

func GenerateToken(qid string, uid string, rid string, ttl int64) (string, error) {
	expirationTime := time.Now().Add(time.Duration(ttl) * time.Minute)
	claims := &Claims{
		UserID:     uid,
		QuestionID: qid,
		RequestID:  rid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: (expirationTime.Unix()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
