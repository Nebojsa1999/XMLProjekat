package config

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/api"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/middleware"
	"github.com/gorilla/mux"
)

type Handlers struct {
	UserHandler                       *api.UserHandler
	CompanyHandler                    *api.CompanyHandler
	CompanyRegistrationRequestHandler *api.CompanyRegistrationRequestHandler
	JobHandler                        *api.JobHandler
	CommentHandler                    *api.CommentHandler
	WageHandler                       *api.WageHandler
	InterviewHandler                  *api.InterviewHandler
}

func ConfigureRouter(handlers Handlers) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/agent-app/user/{id:[0-9a-f]+}",
		handlers.UserHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/user",
		handlers.UserHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/user/register",
		handlers.UserHandler.RegisterANewUser).Methods("POST")
	router.HandleFunc("/agent-app/user/login",
		handlers.UserHandler.Login).Methods("POST")
	router.HandleFunc("/agent-app/user/{id:[0-9a-f]+}",
		handlers.UserHandler.Update).Methods("PUT")

	router.HandleFunc("/agent-app/company/{id:[0-9a-f]+}",
		handlers.CompanyHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/company",
		handlers.CompanyHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/company/register",
		handlers.CompanyHandler.RegisterANewCompany).Methods("POST")
	router.HandleFunc("/agent-app/company/{id:[0-9a-f]+}",
		handlers.CompanyHandler.Update).Methods("PUT")

	router.HandleFunc("/agent-app/company-registration-request/{id:[0-9a-f]+}",
		handlers.CompanyRegistrationRequestHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/company-registration-request/pending",
		handlers.CompanyRegistrationRequestHandler.GetPendingOnes).Methods("GET")
	router.HandleFunc("/agent-app/company-registration-request/accepted",
		handlers.CompanyRegistrationRequestHandler.GetAcceptedOnes).Methods("GET")
	router.HandleFunc("/agent-app/company-registration-request/rejected",
		handlers.CompanyRegistrationRequestHandler.GetRejectedOnes).Methods("GET")
	router.HandleFunc("/agent-app/company-registration-request",
		handlers.CompanyRegistrationRequestHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/company-registration-request",
		handlers.CompanyRegistrationRequestHandler.CreateCompanyRegistrationRequest).Methods("POST")
	router.HandleFunc("/agent-app/company-registration-request/{id:[0-9a-f]+}/update-by-owner",
		handlers.CompanyRegistrationRequestHandler.UpdateByOwner).Methods("PUT")
	router.HandleFunc("/agent-app/company-registration-request/{id:[0-9a-f]+}/update-by-administrator",
		handlers.CompanyRegistrationRequestHandler.UpdateByAdministrator).Methods("PUT")

	router.HandleFunc("/agent-app/job/{id:[0-9a-f]+}",
		handlers.JobHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/job",
		handlers.JobHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/job/create",
		handlers.JobHandler.CreateNewJob).Methods("POST")
	router.HandleFunc("/agent-app/job/{id:[0-9a-f]+}",
		handlers.JobHandler.Update).Methods("PUT")
	router.HandleFunc("/agent-app/job/{id:[0-9a-f]+}",
		handlers.JobHandler.UpdateReviews).Methods("PUT")

	router.HandleFunc("/agent-app/job/comment/{id:[0-9a-f]+}",
		handlers.CommentHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/job/comment/{companyId:[0-9a-f]+}",
		handlers.CommentHandler.GetByCompanyId).Methods("GET")
	router.HandleFunc("/agent-app/job/comment",
		handlers.CommentHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/job/comment/create",
		handlers.CommentHandler.CreateNewComment).Methods("POST")
	router.HandleFunc("/agent-app/job/comment/{id:[0-9a-f]+}",
		handlers.CommentHandler.Update).Methods("PUT")

	router.HandleFunc("/agent-app/job/wage/{id:[0-9a-f]+}",
		handlers.WageHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/job/wage/{companyId:[0-9a-f]+}",
		handlers.WageHandler.GetByCompanyId).Methods("GET")
	router.HandleFunc("/agent-app/job/wage",
		handlers.WageHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/job/wage/create",
		handlers.WageHandler.CreateNewWage).Methods("POST")
	router.HandleFunc("/agent-app/job/wage/{id:[0-9a-f]+}",
		handlers.WageHandler.Update).Methods("PUT")

	router.HandleFunc("/agent-app/job/interview/{id:[0-9a-f]+}",
		handlers.InterviewHandler.Get).Methods("GET")
	router.HandleFunc("/agent-app/job/interview/{companyId:[0-9a-f]+}",
		handlers.InterviewHandler.GetByCompanyId).Methods("GET")
	router.HandleFunc("/agent-app/job/interview",
		handlers.InterviewHandler.GetAll).Methods("GET")
	router.HandleFunc("/agent-app/job/interview/create",
		handlers.InterviewHandler.CreateNewInterview).Methods("POST")
	router.HandleFunc("/agent-app/job/interview/{id:[0-9a-f]+}",
		handlers.InterviewHandler.Update).Methods("PUT")

	router.Use(middleware.IsAuthenticated)

	return router
}
