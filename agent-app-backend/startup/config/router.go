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

	router.Use(middleware.IsAuthenticated)

	return router
}
