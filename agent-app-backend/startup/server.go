package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/persistence"
	cfg "github.com/Nebojsa1999/XMLProjekat/agent-app-backend/startup/config"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
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
	companyRegistrationRequestStore := server.initCompanyRegistrationRequestStore(mongoClient)
	jobStore := server.initJobStore(mongoClient)
	commentStore := server.initCommentStore(mongoClient)
	wageStore := server.initWageStore(mongoClient)
	interviewStore := server.initInterviewStore(mongoClient)

	userService := server.initUserService(userStore)
	companyService := server.initCompanyService(companyStore)
	companyRegistrationRequestService := server.initCompanyRegistrationRequestService(companyRegistrationRequestStore)
	jobService := server.initJobService(jobStore)
	commentService := server.initCommentService(commentStore)
	wageService := server.initWageService(wageStore)
	interviewService := server.initInterviewService(interviewStore)

	userHandler := server.initUserHandler(userService)
	companyHandler := server.initCompanyHandler(companyService)
	companyRegistrationRequestHandler := server.initCompanyRegistrationRequestHandler(companyRegistrationRequestService)
	jobHandler := server.initJobHandler(jobService)
	commentHandler := server.initCommentHandler(commentService)
	wageHandler := server.initWageHandler(wageService)
	interviewHandler := server.initInterviewHandler(interviewService)

	server.startHttpServer(cfg.Handlers{
		UserHandler:                       userHandler,
		CompanyHandler:                    companyHandler,
		CompanyRegistrationRequestHandler: companyRegistrationRequestHandler,
		JobHandler:                        jobHandler,
		CommentHandler:                    commentHandler,
		WageHandler:                       wageHandler,
		InterviewHandler:                  interviewHandler,
	})
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

func (server *Server) initCompanyRegistrationRequestStore(client *mongo.Client) domain.CompanyRegistrationRequestStore {
	store := persistence.NewCompanyRegistrationRequestMongoDBStore(client)
	_, err := store.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, request := range companyRegistrationRequests {
		_, err := store.CreateCompanyRegistrationRequest(request)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initJobStore(client *mongo.Client) domain.JobStore {
	store := persistence.NewJobMongoDBStore(client)
	_, err := store.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, job := range jobs {
		_, err := store.CreateNewJob(job)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initCommentStore(client *mongo.Client) domain.CommentStore {
	store := persistence.NewCommentMongoDBStore(client)
	_, err := store.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, comment := range comments {
		_, err := store.CreateNewComment(comment)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initWageStore(client *mongo.Client) domain.WageStore {
	store := persistence.NewWageMongoDBStore(client)
	_, err := store.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, wage := range wages {
		_, err := store.CreateNewWage(wage)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initInterviewStore(client *mongo.Client) domain.InterviewStore {
	store := persistence.NewInterviewMongoDBStore(client)
	_, err := store.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, interview := range interviews {
		_, err := store.CreateNewInterview(interview)
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

func (server *Server) initCompanyRegistrationRequestService(store domain.CompanyRegistrationRequestStore) *application.CompanyRegistrationRequestService {
	return application.NewCompanyRegistrationRequestService(store)
}

func (server *Server) initJobService(store domain.JobStore) *application.JobService {
	return application.NewJobService(store)
}

func (server *Server) initCommentService(store domain.CommentStore) *application.CommentService {
	return application.NewCommentService(store)
}

func (server *Server) initWageService(store domain.WageStore) *application.WageService {
	return application.NewWageService(store)
}

func (server *Server) initInterviewService(store domain.InterviewStore) *application.InterviewService {
	return application.NewInterviewService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) initCompanyHandler(service *application.CompanyService) *api.CompanyHandler {
	return api.NewCompanyHandler(service)
}

func (server *Server) initCompanyRegistrationRequestHandler(service *application.CompanyRegistrationRequestService) *api.CompanyRegistrationRequestHandler {
	return api.NewCompanyRegistrationRequestHandler(service)
}

func (server *Server) initJobHandler(service *application.JobService) *api.JobHandler {
	return api.NewJobHandler(service)
}

func (server *Server) initCommentHandler(service *application.CommentService) *api.CommentHandler {
	return api.NewCommentHandler(service)
}

func (server *Server) initWageHandler(service *application.WageService) *api.WageHandler {
	return api.NewWageHandler(service)
}

func (server *Server) initInterviewHandler(service *application.InterviewService) *api.InterviewHandler {
	return api.NewInterviewHandler(service)
}

func (server *Server) startHttpServer(handlers cfg.Handlers) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := cfg.ConfigureRouter(handlers)

	corsSpecification := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200", "http://localhost:4201"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler := corsSpecification.Handler(router)

	httpServer := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", server.config.Port), Handler: handler}
	go func() {
		log.Println("agent_app http server starting")

		if err := httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			} else {
				log.Println("agent_app http server was closed due to error")
			}
		}
	}()

	<-quit

	log.Println("agent_app http server shutting down...")

	ctx, cancelFunction := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancelFunction()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("agent_app http server stopped")
}
