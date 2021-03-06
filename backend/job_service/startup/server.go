package startup

import (
	"fmt"
	job "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/job_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/job_service/application"
	"github.com/Nebojsa1999/XMLProjekat/backend/job_service/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/job_service/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/backend/job_service/infrastructure/persistence"
	cfg "github.com/Nebojsa1999/XMLProjekat/backend/job_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *cfg.Config
}

func NewServer(config *cfg.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	jobStore := server.initJobStore(mongoClient)

	jobService := server.initJobService(jobStore)

	jobHandler := server.initJobHandler(jobService)

	server.startGrpcServer(jobHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.JobDBHost, server.config.JobDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initJobStore(client *mongo.Client) domain.JobStore {
	store := persistence.NewJobMongoDBStore(client)
	store.DeleteAll()
	for _, job := range jobs {
		_, err := store.Insert(job)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initJobService(store domain.JobStore) *application.JobService {
	return application.NewJobService(store)
}

func (server *Server) initJobHandler(service *application.JobService) *api.JobHandler {
	return api.NewJobHandler(service)
}

func (server *Server) startGrpcServer(jobHandler *api.JobHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	job.RegisterJobServiceServer(grpcServer, jobHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
