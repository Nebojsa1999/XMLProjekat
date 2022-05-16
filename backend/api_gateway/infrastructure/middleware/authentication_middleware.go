package middleware

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func IsAuthenticated(handler *runtime.ServeMux) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if isAProtectedRoute(request.Method, request.URL.Path) {
			authorizationHeader := request.Header["Authorization"]
			if authorizationHeader != nil {
				tokenWithoutBearerPrefix := strings.TrimPrefix(authorizationHeader[0], "Bearer ")

				token, err := jwt.Parse(tokenWithoutBearerPrefix, func(token *jwt.Token) (interface{}, error) {
					if _, isSigningMethodValid := token.Method.(*jwt.SigningMethodHMAC); !isSigningMethodValid {
						return nil, fmt.Errorf("invalid signing method")
					}

					audience := "billing.jwtgo.io"
					isAudienceValid := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
					if !isAudienceValid {
						return nil, fmt.Errorf("invalid audience")
					}

					issuer := "jwtgo.io"
					isIssuerValid := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
					if !isIssuerValid {
						return nil, fmt.Errorf("invalid issuer")
					}

					return []byte(os.Getenv("SECRET_FOR_JWT")), nil
				})

				if err != nil {
					fmt.Fprintf(writer, err.Error())
					return
				}

				if token.Valid {
					handler.ServeHTTP(writer, request)
					return
				}
			} else {
				fmt.Fprintf(writer, "Authorization token has not been provided!")
				return
			}
		}

		handler.ServeHTTP(writer, request)
	}
}

func isAProtectedRoute(method, path string) bool {
	isPathToPostsOfPublicUser, _ := regexp.MatchString("/user/[0-9a-f]{24}/public", path)

	switch method {
	case "GET":
		if isPathToPostsOfPublicUser || path == "/post/public" {
			return false
		}
	case "POST":
		if path == "/user/register" || path == "/user/login" {
			return false
		}
	}

	return true
}
