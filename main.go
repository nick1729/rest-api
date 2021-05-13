package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Open DB and check connection
func init() {

	var (
		c       tConfig
		connStr string
		err     error
	)

	c, err = loadCfg("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Ip, c.Port, c.Login, c.Pass, c.Table)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
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
