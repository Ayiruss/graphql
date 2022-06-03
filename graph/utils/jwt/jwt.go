package jwt

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Ayiruss/bookstore/graph/utils/errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret = []byte("secret")
)

// JwtCustomClaim strutucte to stores the additional custom element Username apart from the standardclaims
type JwtCustomClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Generates the secure token on login
func GenerateToken(ctx context.Context, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error Generating Key",
		}
	}
	return tokenString, nil
}

func ParseToken(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}
