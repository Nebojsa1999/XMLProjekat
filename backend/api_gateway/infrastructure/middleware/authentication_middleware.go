package middleware

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Role string

const (
	UndefinedRole Role = ""
	CommonUser         = "CommonUser"
	Administrator      = "Administrator"
)

type AuthorizationDeterminingData struct {
	UserId   string
	UserRole string
	Method   string
	Path     string
}

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
					http.Error(writer, err.Error(), http.StatusExpectationFailed)
					return
				}

				if token.Valid {
					authorizationDeterminingData := AuthorizationDeterminingData{
						UserId:   token.Claims.(jwt.MapClaims)["id"].(string),
						UserRole: token.Claims.(jwt.MapClaims)["role"].(string),
						Method:   request.Method,
						Path:     request.URL.Path,
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
	}
}

func isAProtectedRoute(method, path string) bool {
	isPathToPostsOfPublicUser, _ := regexp.MatchString("/user/[0-9a-f]{24}/public", path)

	isPathToConnection, _ := regexp.MatchString("/connection/[0-9a-f]+", path)
	pathToConnections := "/connection"

	isPathToJob, _ := regexp.MatchString("/job/[0-9a-f]+", path)
	pathToAllJobs := "/job/jobs"
	pathToJobAdding := "/job"
	pathToJobEdit := "/job/editJob"

	pathToCreationOfPostFromAgentAppByAPIToken := "/post/job-offer"
	isPathToPostFromUser, _ := regexp.MatchString("/user/[0-9a-f]+/post/[0-9a-f]+", path)
	isPathToAllPostsFromUser, _ := regexp.MatchString("/user/[0-9a-f]+/post", path)
	isPathToCommentsOfPost, _ := regexp.MatchString("/user/[0-9a-f]+/post/[0-9a-f]+/comment", path)
	isPathToLikeOrDislikeOfPost, _ :=
		regexp.MatchString("/user/[0-9a-f]+/post/[0-9a-f]+/liked_or_disliked_by/(dis)?like", path)

	isPathToUser, _ := regexp.MatchString("/user/[0-9a-f]+", path)
	pathToAllUsers := "/user"
	isPathToSearchOfJobsByUser, _ := regexp.MatchString("/job/searchByUser/[0-9a-f]+", path)
	isPathToSearchOfJobsByDescription, _ :=
		regexp.MatchString("/job/searchByDescription/[0-9a-f]+", path)
	isPathToSearchOfJobsByPosition, _ :=
		regexp.MatchString("/job/searchByPosition/[0-9a-f]+", path)
	isPathToSearchOfJobsByRequirements, _ :=
		regexp.MatchString("/job/searchByRequirements/[0-9a-f]+", path)
	isPathToCheckOfPrivacyOfUser, _ := regexp.MatchString("/user/[0-9a-f]+/is-private", path)
	pathToIdsOfAllPublicUsers := "/user/ids-of-all-public-users"
	isPathToGenerationOfJobOffersAPIToken, _ :=
		regexp.MatchString("/user/[0-9a-f]+/generate-job-offers-api-token", path)

	switch method {
	case http.MethodGet:
		if isPathToPostsOfPublicUser || isPathToConnection || path == "/post/public" || path == "/user/public" || isPathToJob ||
			path == pathToAllJobs || isPathToSearchOfJobsByUser || isPathToSearchOfJobsByDescription ||
			isPathToSearchOfJobsByPosition || isPathToSearchOfJobsByRequirements || isPathToPostFromUser ||
			isPathToAllPostsFromUser || isPathToCommentsOfPost || isPathToUser || path == pathToAllUsers ||
			isPathToCheckOfPrivacyOfUser || path == pathToIdsOfAllPublicUsers ||
			isPathToGenerationOfJobOffersAPIToken {
			return false
		}
	case http.MethodPost:
		if path == pathToCreationOfPostFromAgentAppByAPIToken || path == pathToConnections ||
			path == pathToJobAdding || path == "/user/register" || path == "/user/login" ||
			path == "/user/search" || isPathToCommentsOfPost {
			return false
		}
	case http.MethodPut:
		if path == pathToConnections || path == pathToJobEdit || isPathToLikeOrDislikeOfPost || isPathToUser {
			return false
		}
	case http.MethodDelete:
		if strings.HasPrefix(path, "/connection") {
			return false
		}
	}

	return true
}

func isUserAuthorizedToAccessRoute(authorizationDeterminingData AuthorizationDeterminingData) bool {
	userId := authorizationDeterminingData.UserId
	method := authorizationDeterminingData.Method
	path := authorizationDeterminingData.Path

	pathToAllPostsFromUser, _ := regexp.MatchString("/user/[0-9a-f]+/post", path)

	if pathToAllPostsFromUser && method == http.MethodPost {
		userIdInPath := strings.TrimPrefix(path, "/user/")
		userIdInPath = strings.TrimSuffix(userIdInPath, "/post")

		if userId == userIdInPath {
			return true
		}

		return false
	}

	return false
}
