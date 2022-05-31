package middleware

import (
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type AuthorizationDeterminingData struct {
	UserId                 string
	UserRole               enums.UserRole
	OwnedCompanyId         string
	IssuedCompanyRequestId string
	Method                 string
	Path                   string
}

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
					authorizationDeterminingData := AuthorizationDeterminingData{
						UserId:                 token.Claims.(jwt.MapClaims)["id"].(string),
						UserRole:               token.Claims.(jwt.MapClaims)["role"].(enums.UserRole),
						OwnedCompanyId:         token.Claims.(jwt.MapClaims)["ownedCompanyId"].(string),
						IssuedCompanyRequestId: token.Claims.(jwt.MapClaims)["issuedCompanyRequestId"].(string),
						Method:                 request.Method,
						Path:                   request.URL.Path,
					}

					if !isUserAuthorizedToAccessRoute(authorizationDeterminingData) {
						//writer.WriteHeader(http.StatusForbidden)
						fmt.Fprintf(writer, "You are not authorized to access this functionality!")
						return
					}

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
	pathToSingleCompanyById, _ := regexp.MatchString("/agent-app/company/[0-9a-f]+", path)
	pathToAllCompanies := "/agent-app/company"

	pathToUserRegistration := "/agent-app/user/register"
	pathToUserLogin := "/agent-app/user/login"

	switch method {
	case "GET":
		if pathToSingleCompanyById || path == pathToAllCompanies {
			return false
		}
	case "POST":
		if path == pathToUserRegistration || path == pathToUserLogin {
			return false
		}
	}

	return true
}

func isUserAuthorizedToAccessRoute(authorizationDeterminingData AuthorizationDeterminingData) bool {
	userId := authorizationDeterminingData.UserId
	userRole := authorizationDeterminingData.UserRole
	ownedCompanyId := authorizationDeterminingData.OwnedCompanyId
	issuedCompanyRequestId := authorizationDeterminingData.IssuedCompanyRequestId
	method := authorizationDeterminingData.Method
	path := authorizationDeterminingData.Path

	pathToUser, _ := regexp.MatchString("/agent-app/user/[0-9a-f]+", path)
	pathToAllUsers := "/agent-app/user"

	pathToCompany, _ :=
		regexp.MatchString("/agent-app/company/[0-9a-f]+", path)
	pathToCompanyRegistration := "/agent-app/company/register"

	pathToCompanyRegistrationRequest, _ :=
		regexp.MatchString("/agent-app/company-registration-request/[0-9a-f]+", path)
	pathToPendingCompanyRegistrationRequests := "/agent-app/company-registration-request/pending"
	pathToAcceptedCompanyRegistrationRequests := "/agent-app/company-registration-request/accepted"
	pathToRejectedCompanyRegistrationRequests := "/agent-app/company-registration-request/rejected"
	pathToAllCompanyRegistrationRequests := "/agent-app/company-registration-request"
	pathToCompanyRegistrationRequestUpdateByOwner, _ :=
		regexp.MatchString("/agent-app/company-registration-request/[0-9a-f]+/update-by-owner", path)
	pathToCompanyRegistrationRequestUpdateByAdministrator, _ :=
		regexp.MatchString("/agent-app/company-registration-request/[0-9a-f]+/update-by-administrator", path)

	if pathToUser && method == "GET" {
		if userRole == enums.Administrator {
			return true
		}

		idInPath := strings.TrimPrefix(path, "/agent-app/user/")
		if userId == idInPath {
			return true
		}

		return false
	}

	if path == pathToAllUsers && method == "GET" {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	if pathToUser && method == "PUT" {
		if userRole == enums.Administrator {
			return true
		}

		idInPath := strings.TrimPrefix(path, "/agent-app/user/")
		if userId == idInPath {
			return true
		}

		return false
	}

	if path == pathToCompanyRegistration && method == "POST" {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	if pathToCompany && method == "PUT" {
		if userRole == enums.CompanyOwner {
			idInPath := strings.TrimPrefix(path, "/agent-app/company/")
			if ownedCompanyId == idInPath {
				return true
			}
		}

		return false
	}

	if (pathToCompanyRegistrationRequest || path == pathToPendingCompanyRegistrationRequests ||
		path == pathToAcceptedCompanyRegistrationRequests ||
		path == pathToRejectedCompanyRegistrationRequests ||
		path == pathToAllCompanyRegistrationRequests) && method == "GET" {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	if path == pathToAllCompanyRegistrationRequests && method == "POST" {
		if userRole == enums.CommonUser {
			return true
		}

		return false
	}

	if pathToCompanyRegistrationRequestUpdateByOwner && method == "PUT" {
		if userRole == enums.CompanyOwner {
			idInPath := strings.TrimPrefix(path, "/agent-app/company-registration-request/")
			if issuedCompanyRequestId == idInPath {
				return true
			}
		}

		return false
	}

	if pathToCompanyRegistrationRequestUpdateByAdministrator && method == "PUT" {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	return false
}
