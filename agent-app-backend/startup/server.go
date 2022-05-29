package startup

import (
	"context"
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/persistence"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
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

	server.startHttpServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AgentAppDBHost, server.config.AgentAppDBPort)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	_, err := store.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
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

func (server *Server) startHttpServer(userHandler *api.UserHandler) {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/agent-app/user", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/user", userHandler.RegisterANewUser).Methods("POST")
	router.HandleFunc("/agent-app/user/{id:[0-9a-z]+}", userHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/user/{id:[0-9a-z]+}", userHandler.Update).Methods("PUT")

	httpServer := &http.Server{Addr: fmt.Sprintf(":%s", server.config.Port), Handler: router}
	go func() {
		log.Println("agent_app http server starting")

		if err := httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	log.Println("agent_app http server shutting down...")

	ctx, cancelFunction := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancelFunction()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("agent_app http server stopped")
}
