package service

import (
	"context"
	"net/http"

	"github.com/Ayiruss/bookstore/graph/internal/domain/users"
	"github.com/Ayiruss/bookstore/graph/utils/errors"
	"github.com/Ayiruss/bookstore/graph/utils/jwt"
)

type authString string

// Authenticate authenticates the valid user and store the custom claims in the context
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := jwt.ParseToken(context.Background(), auth)
		if err != nil || !validate.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		customClaim, _ := validate.Claims.(*jwt.JwtCustomClaim)

		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) *jwt.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*jwt.JwtCustomClaim)
	return raw
}

// Login verifies the user in the database and creates a secure token for further transactions
func Login(ctx context.Context, username string, password string) (interface{}, error) {
	user := users.User{
		Username: username,
	}
	err := user.GetByUserName()
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusBadRequest,
			Message:    "User not found",
		}
	}

	if err := user.CheckPasswordHash(password); err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Username or Password",
		}
	}
	token, err := jwt.GenerateToken(ctx, username)
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to generate the token for the given combination",
		}
	}

	return map[string]interface{}{
		"token": token,
	}, nil

}
