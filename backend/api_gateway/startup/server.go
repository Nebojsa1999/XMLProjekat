package startup

import (
	"context"
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/middleware"
	cfg "github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/startup/config"
	postingGw "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	userGw "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	server.initHandlers()
	server.initCustomHandlers()

	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		log.Fatalf(err.Error())
	}

	postingEndpoint := fmt.Sprintf("%s:%s", server.config.PostingHost, server.config.PostingPort)
	err = postingGw.RegisterPostingServiceHandlerFromEndpoint(context.TODO(), server.mux, postingEndpoint, opts)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postingEndpoint := fmt.Sprintf("%s:%s", server.config.PostingHost, server.config.PostingPort)

	registerHandler := api.NewRegisterHandler(userEndpoint)
	registerHandler.Init(server.mux)

	publicPostHandler := api.NewPublicPostHandler(userEndpoint, postingEndpoint)
	publicPostHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), middleware.IsAuthenticated(server.mux)))
}
