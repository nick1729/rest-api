package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dUser = "login"
	dPass = "pass"
	dHost = "url"
	dPort = "port"
	dName = "table"
)

var db *sql.DB

// Open DB and check connection
func init() {

	var (
		connStr         string
		errCon, errPing error
	)

	connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dHost, dPort, dUser, dPass, dName)

	db, errCon = sql.Open("postgres", connStr)
	if errCon != nil {
		log.Fatal(errCon)
	}

	errPing = db.Ping()
	if errPing != nil {
		log.Fatal(errPing)
	}
}

//Show start page
func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Its start page! If u can read this db works correctly =)")
}

func main() {

	//Routing
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
