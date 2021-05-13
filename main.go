package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

// Show start page
func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Its start page! If u can read this db works correctly =)")
}

// Add new user
// http://127.0.0.1:8000/users?Firstname=Gena&Lastname=Ivanov&Email=gena@mail.gg&Age=44
func addUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var (
		user tUser
		err  error
		msg  string
	)

	user.Firstname = r.FormValue("Firstname")
	user.Lastname = r.FormValue("Lastname")
	user.Email = r.FormValue("Email")
	user.Age, err = strconv.ParseUint(r.FormValue("Age"), 10, 32)

	switch {
	case user.Firstname == "" || len(user.Firstname) > 256:
		http.Error(w, "Incorrect users firstname!", http.StatusBadRequest)
		return
	case user.Lastname == "" || len(user.Lastname) > 256:
		http.Error(w, "Incorrect users lastname!", http.StatusBadRequest)
		return
	case user.Email == "" || len(user.Email) > 256 || isEmail(user.Email) != true:
		http.Error(w, "Incorrect users email!", http.StatusBadRequest)
		return
	case err != nil || user.Age < 1 || user.Age > 256:
		http.Error(w, "Incorrect users age!", http.StatusBadRequest)
		return
	}

	db.QueryRow("insert into Users (firstname, lastname, email, age, created) values ($1, $2, $3, $4, $5) returning id",
		user.Firstname, user.Lastname, user.Email, user.Age, time.Now()).Scan(&user.ID)

	msg = fmt.Sprintf("The user %v was added", user.ID)
	fmt.Fprint(w, msg)
}

// Show user by ID
func showUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	vars := mux.Vars(r)
	id := vars["key"]

	if isUUID(id) != true {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	user := tUser{}

	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname,
		&user.Email, &user.Age, &user.Created)
	switch err {
	case sql.ErrNoRows:
		http.Error(w, "Not found", http.StatusNotFound)
		return
	case nil:
		log.Print("Request completed successfully")
	default:
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("User data:\n%v, %s, %s, %d, %v\n", user.ID,
		user.Firstname, user.Lastname, user.Age, user.Created)
	fmt.Fprint(w, msg)
}

func main() {

	//Routing
	router := mux.NewRouter()
	router.HandleFunc("/users/{key}", showUser)
	router.HandleFunc("/users", addUser)
	router.HandleFunc("/", homePage)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
