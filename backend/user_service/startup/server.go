package startup

import (
	"fmt"
	user "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/application"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/infrastructure/persistence"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	userService := server.initUserService(userStore)

	userHandler := server.initUserHandler(userService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll()
	for _, user := range users{
		_, err := store.RegisterANewUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
