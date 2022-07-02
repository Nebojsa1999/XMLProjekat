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
	isPathToUserProfile, _ := regexp.MatchString("/user/[0-9a-f]{24}", path)
	pathToJob, _ := regexp.MatchString("/job/[0-9a-f]+", path)
	pathToAllJobs := "/job/jobs"
	pathToJobAdding := "/job"
	switch method {
	case "GET":
		if isPathToPostsOfPublicUser || path == "/post/public" || path == "/user/public" || isPathToUserProfile || pathToJob || path == pathToAllJobs {
			return false
		}
	case "POST":
		if path == pathToJobAdding || path == "/user/register" || path == "/user/login" || path == "/user/search" {
			return false
		}
	}

	return true
}

func isUserAuthorizedToAccessRoute(authorizationDeterminingData AuthorizationDeterminingData) bool {
	userId := authorizationDeterminingData.UserId
	userRole := authorizationDeterminingData.UserRole
	method := authorizationDeterminingData.Method
	path := authorizationDeterminingData.Path

	pathToSearchOfJobsByUser, _ := regexp.MatchString("/job/searchByUser/[0-9a-f]+", path)
	pathToSearchOfJobsByDescription, _ :=
		regexp.MatchString("/job/searchByDescription/[0-9a-f]+", path)
	pathToSearchOfJobsByPosition, _ :=
		regexp.MatchString("/job/searchByPosition/[0-9a-f]+", path)
	pathToSearchOfJobsByRequirements, _ :=
		regexp.MatchString("/job/searchByRequirements/[0-9a-f]+", path)
	pathToJobEdit := "/job/editJob"

	pathToPostFromUser, _ := regexp.MatchString("/user/[0-9a-f]+/post/[0-9a-f]+", path)
	pathToAllPostsFromUser, _ := regexp.MatchString("/user/[0-9a-f]+/post", path)
	pathToCommentsOfPost, _ := regexp.MatchString("/user/[0-9a-f]+/post/[0-9a-f]+/comment", path)
	pathToLikeOrDislikeOfPost, _ :=
		regexp.MatchString("/user/[0-9a-f]+/post/[0-9a-f]+/liked_or_disliked_by/(dis)?like", path)

	pathToUser, _ := regexp.MatchString("/user/[0-9a-f]+", path)
	pathToAllUsers := "/user"
	pathToCheckOfPrivacyOfUser, _ := regexp.MatchString("/user/[0-9a-f]+/is-private", path)
	pathToIdsOfAllPublicUsers := "/user/ids-of-all-public-users"
	pathToGenerationOfJobOffersAPIToken, _ :=
		regexp.MatchString("/user/[0-9a-f]+/generate-job-offers-api-token", path)

	if (pathToSearchOfJobsByUser || pathToSearchOfJobsByDescription || pathToSearchOfJobsByPosition ||
		pathToSearchOfJobsByRequirements) && method == http.MethodGet {
		return true
	}

	if path == pathToJobEdit && method == http.MethodPut {
		if userRole == Administrator {
			return true
		}

		return false
	}

	if pathToPostFromUser && method == http.MethodGet {
		return true
	}

	if pathToAllPostsFromUser && method == http.MethodGet {
		return true
	}

	if pathToAllPostsFromUser && method == http.MethodPost {
		userIdInPath := strings.TrimPrefix(path, "/user/")
		userIdInPath = strings.TrimSuffix(userIdInPath, "/post")

		if userId == userIdInPath {
			return true
		}

		return false
	}

	if pathToCommentsOfPost && method == http.MethodGet {
		return true
	}

	if pathToCommentsOfPost && method == http.MethodPost {
		return true
	}

	if pathToLikeOrDislikeOfPost && method == http.MethodPut {
		return true
	}

	if pathToUser && method == http.MethodGet {
		return true
	}

	if pathToUser && method == http.MethodPut {
		if userRole == Administrator {
			return true
		}

		idInPath := strings.TrimPrefix(path, "/user/")
		if userId == idInPath {
			return true
		}

		return false
	}

	if path == pathToAllUsers && method == http.MethodGet {
		return true
	}

	if pathToCheckOfPrivacyOfUser && method == http.MethodGet {
		return true
	}

	if path == pathToIdsOfAllPublicUsers && method == http.MethodGet {
		return true
	}

	if pathToGenerationOfJobOffersAPIToken && method == http.MethodGet {
		return true
	}

	return false
}
