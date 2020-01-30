package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/confus1on/meetmeup/models"
	"github.com/confus1on/meetmeup/postgres"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(repos postgres.UsersRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := repos.GetUserById(context.Background(), claims["jti"].(int))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripeBearerPrefixFromToken,
}

func stripeBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

var MySecretJWT = "mysecretauth"

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(MySecretJWT)
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken Error : ")
}

func GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New("no user in context")
	}

	user, ok := ctx.Value(CurrentUserKey).(*models.User)
	if !ok || &user.ID == nil {
		return nil, errors.New("no user in context")
	}

	return user, nil
}
