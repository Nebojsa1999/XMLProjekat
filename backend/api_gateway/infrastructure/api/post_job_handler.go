package api

import (
	"context"
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/services"
	jobPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/job_service"
	postingPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	userPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"os"
)

type PostJobHandler struct {
	userClientAddress    string
	jobClientAddress     string
	postingClientAddress string
}

func NewPostJobHandler(userClientAddress, jobClientAddress, postingClientAddress string) Handler {
	return &PostJobHandler{
		userClientAddress:    userClientAddress,
		jobClientAddress:     jobClientAddress,
		postingClientAddress: postingClientAddress,
	}
}

func (handler *PostJobHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/post/job-offer", handler.PostAJobOffer)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (handler *PostJobHandler) PostAJobOffer(writer http.ResponseWriter, request *http.Request, pathParams map[string]string) {
	if !isContentTypeValid(writer, request) {
		return
	}

	postJobOfferRequest, err := decodePostJobOfferRequestFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := handler.getDislinktUserIdClaimFromJobOffersAPIToken(postJobOfferRequest.JobOffersAPIToken)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jobOffersAPITokenStatusRequest := &domain.JobOffersAPITokenStatusRequest{UserId: userId}
	err = handler.hasUserGeneratedJobOffersAPIToken(jobOffersAPITokenStatusRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	if !jobOffersAPITokenStatusRequest.HasGeneratedToken {
		http.Error(writer, "User did not generate API token.", http.StatusBadRequest)
		return
	}

	err = handler.createJob(postJobOfferRequest.Job, userId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.createPostWithJobOffer(postJobOfferRequest.Job, userId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	renderJSON(writer, "")
}

func (handler *PostJobHandler) getDislinktUserIdClaimFromJobOffersAPIToken(jobOffersAPIToken string) (string, error) {
	token, err := jwt.Parse(jobOffersAPIToken, func(token *jwt.Token) (interface{}, error) {
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

		return []byte(os.Getenv("SECRET_FOR_JOB_OFFERS_API_TOKEN")), nil
	})

	if err != nil {
		return "", err
	}

	if token.Valid {
		return token.Claims.(jwt.MapClaims)["dislinktUserId"].(string), nil
	}

	return "", fmt.Errorf("job offers api token is not valid")
}

func (handler *PostJobHandler) hasUserGeneratedJobOffersAPIToken(jobOffersAPITokenStatusRequest *domain.JobOffersAPITokenStatusRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)

	userResponse, err := userClient.Get(context.TODO(), &userPb.GetRequest{Id: jobOffersAPITokenStatusRequest.UserId})
	if err != nil {
		return err
	}

	if userResponse.User.JobOffersAPIToken != "" {
		jobOffersAPITokenStatusRequest.HasGeneratedToken = true
	} else {
		jobOffersAPITokenStatusRequest.HasGeneratedToken = false
	}

	return nil
}

func (handler *PostJobHandler) createJob(job *domain.Job, userId string) error {
	jobClient := services.NewJobClient(handler.jobClientAddress)

	jobProtobufObject := jobPb.Job{
		Id:           job.Id,
		UserId:       userId,
		CreatedAt:    timestamppb.New(job.CreatedAt),
		Position:     job.Position,
		Description:  job.Description,
		Requirements: job.Requirements,
	}

	addResponse, err := jobClient.Add(context.TODO(), &jobPb.AddRequest{Job: &jobProtobufObject})
	if err != nil {
		return err
	} else if addResponse.Success != "success" {
		return fmt.Errorf("database already contains job with given id")
	}

	return nil
}

func (handler *PostJobHandler) createPostWithJobOffer(job *domain.Job, userId string) error {
	postingClient := services.NewPostingClient(handler.postingClientAddress)

	postWithJobOffer := postingPb.Post{
		Id:            "",
		OwnerId:       userId,
		Content:       "Nova ponuda za posao:",
		Image:         "",
		LikesCount:    0,
		DislikesCount: 0,
		Comments:      nil,
		Link:          []string{"http://localhost:8001/agent-app/job/" + job.Id},
		WhoLiked:      nil,
		WhoDisliked:   nil,
		PostedAt:      timestamppb.Now().String(),
	}

	_, err := postingClient.CreatePost(context.TODO(), &postingPb.NewPostRequest{
		Id:   userId,
		Post: &postWithJobOffer,
	})
	if err != nil {
		return err
	}

	return nil
}
