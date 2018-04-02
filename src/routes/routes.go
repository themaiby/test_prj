package routes

import (
	"github.com/gorilla/mux"
	log "github.com/kataras/golog"

	"../controller"
	"../middleware"
)

func Make() *mux.Router {
	log.Info("Registering routes")
	routes := mux.NewRouter()

	routes.Use(middleware.SetRequestTime)

	// TODO: auth middleware

	/*
		User
	*/

	// domain.com/users
	routes.HandleFunc("/users/", controller.UserController.GetAllUsers).
		Methods("GET").
		Name("UserGetterAll")

	// domain.com/users/24
	routes.HandleFunc("/users/{id:[0-9]+}", controller.UserController.GetUser).
		Methods("GET").
		Name("UserGetter")

	/*
		Auth
	*/

	// domain.com/login TODO: ldap auth
	routes.HandleFunc("/login", controller.CapController.Cap).
		Methods("POST").
		Name("Login")

	/*
		Mailing
	*/

	// domain.com/mail/75
	routes.HandleFunc("/mail/{id:[0-9]+}", controller.CapController.Cap).
		Methods("GET").
		Name("EmailGetter")

	// domain.com/mail/new
	routes.HandleFunc("/mail/new", controller.CapController.Cap).
		Methods("POST").
		Name("EmailCreator")

	// domain.com/mail/take/75
	routes.HandleFunc("/mail/take", controller.CapController.Cap).
		Methods("POST").
		Name("EmailTaker")

	return routes
}
