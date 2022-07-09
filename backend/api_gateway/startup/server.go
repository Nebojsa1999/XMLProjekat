package startup

import (
	"context"
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/middleware"
	cfg "github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/startup/config"
	connectionGw "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	jobGw "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/job_service"
	postingGw "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	userGw "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
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

	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	err = jobGw.RegisterJobServiceHandlerFromEndpoint(context.TODO(), server.mux, jobEndpoint, opts)
	if err != nil {
		log.Fatalf(err.Error())
	}

	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	err = connectionGw.RegisterConnectionServiceHandlerFromEndpoint(context.TODO(), server.mux, connectionEndpoint, opts)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postingEndpoint := fmt.Sprintf("%s:%s", server.config.PostingHost, server.config.PostingPort)
	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)

	registerHandler := api.NewRegisterHandler(userEndpoint)
	registerHandler.Init(server.mux)

	publicPostHandler := api.NewPublicPostHandler(userEndpoint, postingEndpoint)
	publicPostHandler.Init(server.mux)

	postJobHandler := api.NewPostJobHandler(userEndpoint, jobEndpoint, postingEndpoint)
	postJobHandler.Init(server.mux)

	postsOfFollowingHandler := api.NewPostsOfFollowingHandler(connectionEndpoint, postingEndpoint)
	postsOfFollowingHandler.Init(server.mux)
}

func (server *Server) Start() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200", "http://localhost:4201"},
		AllowedMethods: []string{http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(middleware.IsAuthenticated(server.mux))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handler))
}
