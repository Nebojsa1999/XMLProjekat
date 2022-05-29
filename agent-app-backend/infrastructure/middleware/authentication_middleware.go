package middleware

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func IsAuthenticated(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
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

					return []byte(os.Getenv("SECRET_FOR_AGENT_APP_JWT")), nil
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
	})
}

func isAProtectedRoute(method, path string) bool {
	pathToSingleCompany, _ := regexp.MatchString("/agent-app/company/[0-9a-f]+", path)

	switch method {
	case "GET":
		if pathToSingleCompany || path == "/agent-app/company" {
			return false
		}
	case "POST":
		if path == "/agent-app/user/register" || path == "/agent-app/user/login" {
			return false
		}
	}

	return true
}
