package startup

import (
	"fmt"
	connection "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/application"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/infrastructure/persistence"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/startup/config"
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
	connectionStore := server.initConnectionStore(mongoClient)

	connectionService := server.initConnectionService(connectionStore)

	connectionHandler := server.initConnectionHandler(connectionService)

	server.startGrpcServer(connectionHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ConnectionDBHost, server.config.ConnectionDBPort)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (server *Server) initConnectionStore(client *mongo.Client) domain.ConnectionStore {
	store := persistence.NewConnectionMongoDBStore(client)
	store.DeleteAll()
	for _, connection := range connections{
		_, err := store.Create(connection)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initConnectionService(store domain.ConnectionStore) *application.ConnectionService {
	return application.NewConnectionService(store)
}

func (server *Server) initConnectionHandler(service *application.ConnectionService) *api.ConnectionHandler {
	return api.NewConnectionHandler(service)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	connection.RegisterConnectionServiceServer(grpcServer, connectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
