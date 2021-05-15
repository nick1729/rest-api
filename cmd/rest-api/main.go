package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/internal/database"
	"rest-api/internal/handlers"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	wait time.Duration
	err  error
)

func init() {

	handlers.DB, err = database.Dial()
	if err != nil {
		log.Panic(err)
	}

	log.Print("DB connected")
}

func main() {

	// Routing
	router := mux.NewRouter()
	router.HandleFunc("/users/", handlers.EditUser).Methods("PUT")
	router.HandleFunc("/users/{key}", handlers.ShowUser).Methods("GET")
	router.HandleFunc("/users", handlers.AddUser).Methods("POST")
	http.Handle("/", router)

	srv := &http.Server{
		Addr:         "localhost:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Start service
	go func() {
		log.Print("Server is listening...")
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	wait = time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Print("Shutting down...")
	os.Exit(0)
}
