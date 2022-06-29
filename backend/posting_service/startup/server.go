package startup

import (
	"fmt"
	"log"
	"net"

	ps "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/application"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/infrastructure/persistence"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
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
	postStore := server.initPostStore(mongoClient)

	postingService := server.initPostingService(postStore)

	postHandler := server.initPostHandler(postingService)

	server.startGrpcServer(postHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostingDBHost, server.config.PostingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initPostStore(client *mongo.Client) domain.PostStore {
	store := persistence.NewPostMongoDBStore(client)
	store.DeleteAll()

	for _, postRequest := range posts {
		id := postRequest.OwnerId
		_, err := store.CreatePost(id, &postRequest)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initPostingService(store domain.PostStore) *application.PostService {
	return application.NewPostService(store)
}

func (server *Server) initPostHandler(service *application.PostService) *api.PostHandler {
	return api.NewPostHandler(service)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	ps.RegisterPostingServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
