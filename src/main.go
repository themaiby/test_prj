package main

import (
	// application internal
	"./configure"
	"./routes"

	// 3th packages
	"./db"
	"./middleware"
	"context"
	"github.com/gorilla/mux"
	log "github.com/kataras/golog"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Setup Log Format, Level
	// Debug can be enabled by -debug flag
	configure.Logging()
	// read variables from config file
	configure.Application()

	//connect DB
	db.Connect()

	// init routing
	routing := routes.Make()

	//	migration.Run() // TODO: make flag -migration

	// run
	run(routing)
}

func run(routing *mux.Router) {
	var domain, port string
	domain = viper.GetString("application.domain")
	port = viper.GetString("application.port")

	// run
	log.Info("Server is running at " + domain + ":" + port)
	log.Println("Press Ctrl+C to shutdown server")

	// http.Handle("/", routing)

	srv := &http.Server{
		Handler: middleware.LogWrapper{http.Handler(routing)},
		Addr:    domain + ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Warn(err)
			return
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	var wait time.Duration
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("Shutting down\n")
	os.Exit(0)
}
