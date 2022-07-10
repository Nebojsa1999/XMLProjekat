package middleware

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthorizationDeterminingData struct {
	UserId                 string
	UserRole               string
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
					http.Error(writer, err.Error(), http.StatusExpectationFailed)
					return
				}

				if token.Valid {
					authorizationDeterminingData := AuthorizationDeterminingData{
						UserId:                 token.Claims.(jwt.MapClaims)["id"].(string),
						UserRole:               token.Claims.(jwt.MapClaims)["role"].(string),
						OwnedCompanyId:         token.Claims.(jwt.MapClaims)["ownedCompanyId"].(string),
						IssuedCompanyRequestId: token.Claims.(jwt.MapClaims)["issuedCompanyRequestId"].(string),
						Method:                 request.Method,
						Path:                   request.URL.Path,
					}

					if !isUserAuthorizedToAccessRoute(authorizationDeterminingData) {
						http.Error(writer, "You are not authorized to access this functionality!", http.StatusUnauthorized)
						return
					}

					handler.ServeHTTP(writer, request)
					return
				}
			} else {
				http.Error(writer, "Authorization token has not been provided!", http.StatusForbidden)
				return
			}
		}

		handler.ServeHTTP(writer, request)
	})
}

func isAProtectedRoute(method, path string) bool {
	pathToSingleCompanyById, _ := regexp.MatchString("/agent-app/company/[0-9a-f]+", path)
	pathToAllCompanies := "/agent-app/company"

	pathToCompanyRegistrationRequestUpdateByOwner, _ :=
		regexp.MatchString("/agent-app/company-registration-request/[0-9a-f]+/update-by-owner", path)

	pathToCompany, _ :=
		regexp.MatchString("/agent-app/company/[0-9a-f]+", path)

	pathToSingleJobById, _ := regexp.MatchString("/agent-app/job/[0-9a-f]+", path)
	pathToAllJobs := "/agent-app/job"

	pathToSingleCommentById, _ := regexp.MatchString("/agent-app/job/comment/[0-9a-f]+", path)
	pathToAllComments, _ := regexp.MatchString("/agent-app/company/comment/[0-9a-f]+", path)
	pathToCommentCreation := "/agent-app/job/comment/create"

	pathToSingleWageById, _ := regexp.MatchString("/agent-app/job/wage/[0-9a-f]+", path)
	pathToAllWages, _ := regexp.MatchString("/agent-app/company/wage/[0-9a-f]+", path)
	pathToWageCreation := "/agent-app/job/wage/create"

	pathToSingleInterviewById, _ := regexp.MatchString("/agent-app/job/interview/[0-9a-f]+", path)
	pathToAllInterviews, _ := regexp.MatchString("/agent-app/company/interview/[0-9a-f]+", path)
	pathToInterviewCreation := "/agent-app/job/interview/create"

	pathToUserRegistration := "/agent-app/user/register"
	pathToUserLogin := "/agent-app/user/login"

	switch method {
	case "GET":
		if pathToSingleCompanyById || path == pathToAllCompanies || pathToSingleJobById || path == pathToAllJobs || pathToSingleCommentById || pathToAllComments || pathToSingleWageById || pathToAllWages || pathToSingleInterviewById || pathToAllInterviews {
			return false
		}
	case "POST":
		if path == pathToUserRegistration || path == pathToUserLogin || path == pathToCommentCreation ||
			path == pathToWageCreation || path == pathToInterviewCreation {
			return false
		}

	case "PUT":
		if pathToCompanyRegistrationRequestUpdateByOwner || pathToCompany {
			return false
		}
	}

	return true
}

func isUserAuthorizedToAccessRoute(authorizationDeterminingData AuthorizationDeterminingData) bool {
	userId := authorizationDeterminingData.UserId
	userRole := authorizationDeterminingData.UserRole
	//	ownedCompanyId := authorizationDeterminingData.OwnedCompanyId
	method := authorizationDeterminingData.Method
	path := authorizationDeterminingData.Path

	pathToUser, _ := regexp.MatchString("/agent-app/user/[0-9a-f]+", path)
	pathToAllUsers := "/agent-app/user"

	//pathToCompany, _ :=
	//	regexp.MatchString("/agent-app/company/[0-9a-f]+", path)
	pathToCompanyRegistration := "/agent-app/company/register"

	pathToCompanyRegistrationRequest, _ :=
		regexp.MatchString("/agent-app/company-registration-request/[0-9a-f]+", path)
	pathToPendingCompanyRegistrationRequests := "/agent-app/company-registration-request/pending"
	pathToAcceptedCompanyRegistrationRequests := "/agent-app/company-registration-request/accepted"
	pathToRejectedCompanyRegistrationRequests := "/agent-app/company-registration-request/rejected"
	pathToAllCompanyRegistrationRequests := "/agent-app/company-registration-request"

	pathToCompanyRegistrationRequestUpdateByAdministrator, _ :=
		regexp.MatchString("/agent-app/company-registration-request/[0-9a-f]+/update-by-administrator", path)

	pathToJob, _ :=
		regexp.MatchString("/agent-app/job/[0-9a-f]+", path)
	pathToJobCreation := "/agent-app/job/create"
	pathToJobUpdate, _ := regexp.MatchString("/agent-app/job/[0-9a-f]+/update", path)
	pathToJobUpdateReviews, _ := regexp.MatchString("/agent-app/job/[0-9a-f]+/update-reviews", path)

	pathToComment, _ :=
		regexp.MatchString("/agent-app/company/comment/[0-9a-f]+", path)

	pathToWage, _ :=
		regexp.MatchString("/agent-app/job/wage/[0-9a-f]+", path)

	pathToInterview, _ :=
		regexp.MatchString("/agent-app/job/interview/[0-9a-f]+", path)

	if pathToUser && method == http.MethodGet {
		if userRole == enums.Administrator {
			return true
		}

		idInPath := strings.TrimPrefix(path, "/agent-app/user/")
		if userId == idInPath {
			return true
		}

		return false
	}

	if path == pathToAllUsers && method == http.MethodGet {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	if pathToUser && method == http.MethodPut {
		if userRole == enums.Administrator {
			return true
		}

		idInPath := strings.TrimPrefix(path, "/agent-app/user/")
		if userId == idInPath {
			return true
		}

		return false
	}

	if path == pathToCompanyRegistration && method == http.MethodPost {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	//if pathToCompany && method == http.MethodPut {
	//	if userRole == enums.CompanyOwner {
	//		idInPath := strings.TrimPrefix(path, "/agent-app/company/")
	//		if ownedCompanyId == idInPath {
	//			return true
	//		}
	//	}
	//
	//	return false
	//}

	if path == pathToJobCreation && method == "POST" {
		if userRole == enums.CompanyOwner {
			return true
		}

		return false
	}

	if pathToJob && method == "PUT" {
		if userRole == enums.CompanyOwner {
			return true
		}

		return false
	}

	if pathToComment && method == "PUT" {
		if userRole == enums.CommonUser {
			return true
		}

		return false
	}

	if pathToWage && method == "PUT" {
		if userRole == enums.CommonUser {
			return true
		}

		return false
	}

	if pathToInterview && method == "PUT" {
		if userRole == enums.CommonUser {
			return true
		}

		return false
	}

	if pathToJobUpdate && method == http.MethodPut {
		if userRole == enums.CompanyOwner {
			return true
		}

		return false
	}

	if pathToJobUpdateReviews && method == http.MethodPut {
		if userRole == enums.CommonUser {
			return true
		}

		return false
	}

	if (pathToCompanyRegistrationRequest || path == pathToPendingCompanyRegistrationRequests ||
		path == pathToAcceptedCompanyRegistrationRequests ||
		path == pathToRejectedCompanyRegistrationRequests ||
		path == pathToAllCompanyRegistrationRequests) && method == http.MethodGet {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	if path == pathToAllCompanyRegistrationRequests && method == http.MethodPost {
		if userRole == enums.CommonUser {
			return true
		}

		return false
	}

	if pathToCompanyRegistrationRequestUpdateByAdministrator && method == http.MethodPut {
		if userRole == enums.Administrator {
			return true
		}

		return false
	}

	return false
}
