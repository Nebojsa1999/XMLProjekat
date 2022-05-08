package middleware

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"os"
)

func IsAuthenticated(handler *runtime.ServeMux) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isAProtectedRoute(request.Method, request.URL.Path) {
			if request.Header["Token"] != nil {
				token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
					if _, isSigningMethodValid := token.Method.(*jwt.SigningMethodHMAC); !isSigningMethodValid {
						return nil, fmt.Errorf("Invalid signing method!")
					}

					audience := "billing.jwtgo.io"
					isAudienceValid := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
					if !isAudienceValid {
						return nil, fmt.Errorf("Invalid audience!")
					}

					issuer := "jwtgo.io"
					isIssuerValid := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
					if !isIssuerValid {
						return nil, fmt.Errorf("Invalid issuer!")
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
	})
}

func isAProtectedRoute(method, path string) bool {
	switch method {
	case "GET":
		if path == "/user/{id}/public" {
			return false
		}
	case "POST":
		if path == "/user/register" || path == "/user/login" {
			return false
		}
	}

	return true
}
