package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/internal/database"
	"rest-api/internal/handlers"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	wait time.Duration
	err  error
)

func run() error {
	// open and check DB
	handlers.DB, err = database.Dial()
	if err != nil {
		return err
	}

	log.Print("DB connected")

	// make migrations
	err = database.RunMigrations()
	if err != nil {
		return err
	}

	log.Print("DB updated")

	// routing
	router := mux.NewRouter()
	router.HandleFunc("/users/delete/{key}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/all", handlers.ShowAllUsers).Methods("GET")
	router.HandleFunc("/users/{key}", handlers.ShowUser).Methods("GET")
	router.HandleFunc("/users/", handlers.EditUser).Methods("PUT")
	router.HandleFunc("/users", handlers.AddUser).Methods("POST")
	router.HandleFunc("/", handlers.ShowHomePage).Methods("GET")
	http.Handle("/", router)

	srv := &http.Server{
		Addr:         os.Getenv("HTTP_ADDR"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// start service
	go func() {
		log.Print("Server is listening...")
		if err = srv.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	wait = time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Print("Shutting down...")
	os.Exit(0)

	return nil
}

func main() {

	var err error

	err = run()
	if err != nil {
		log.Fatal(err)
	}
}
