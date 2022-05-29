package startup

import (
	"context"
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/middleware"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/persistence"
	cfg "github.com/Nebojsa1999/XMLProjekat/agent-app-backend/startup/config"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
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
	userStore := server.initUserStore(mongoClient)
	companyStore := server.initCompanyStore(mongoClient)

	userService := server.initUserService(userStore)
	companyService := server.initCompanyService(companyStore)

	userHandler := server.initUserHandler(userService)
	companyHandler := server.initCompanyHandler(companyService)

	server.startHttpServer(userHandler, companyHandler)
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

func (server *Server) initCompanyStore(client *mongo.Client) domain.CompanyStore {
	store := persistence.NewCompanyMongoDBStore(client)
	_, err := store.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, company := range companies {
		_, err := store.RegisterANewCompany(company)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initCompanyService(store domain.CompanyStore) *application.CompanyService {
	return application.NewCompanyService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) initCompanyHandler(service *application.CompanyService) *api.CompanyHandler {
	return api.NewCompanyHandler(service)
}

func (server *Server) startHttpServer(userHandler *api.UserHandler, companyHandler *api.CompanyHandler) {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/agent-app/user/{id:[0-9a-f]+}", userHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/user", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/user/register", userHandler.RegisterANewUser).Methods("POST")
	router.HandleFunc("/agent-app/user/login", userHandler.Login).Methods("POST")
	router.HandleFunc("/agent-app/user/{id:[0-9a-f]+}", userHandler.Update).Methods("PUT")

	router.HandleFunc("/agent-app/company/{id:[0-9a-f]+}", companyHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/company", companyHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/company/register", companyHandler.RegisterANewCompany).Methods("POST")
	router.HandleFunc("/agent-app/company/{id:[0-9a-f]+}", companyHandler.Update).Methods("PUT")

	router.Use(middleware.IsAuthenticated)

	httpServer := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", server.config.Port), Handler: router}
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
