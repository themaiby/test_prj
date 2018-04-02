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

	// example path
	// routes.HandleFunc("/example/{id:[0-9]+}", controller.ExampleController.Example)

	routes.Use(middleware.SetRequestTime)
	// work routes
	routes.HandleFunc("/users/{id:[0-9]+}", controller.MainContoller.MainPage)

	return routes
}
