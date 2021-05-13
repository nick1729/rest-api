package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Email validation
func isEmail(e string) bool {

	var err error

	_, err = mail.ParseAddress(e)
	if err != nil {
		return false
	}

	return true
}

// Email validation
func isUUID(u string) bool {

	var err error

	_, err = uuid.Parse(u)
	if err != nil {
		return false
	}

	return true
}

// Add new user
// http://127.0.0.1:8000/users?Firstname=Gena&Lastname=Ivanov&Email=gena@mail.gg&Age=44
func addUser(w http.ResponseWriter, r *http.Request) {

	var (
		user tUser
		err  error
		msg  string
	)

	// read keys
	user.Firstname = r.FormValue("Firstname")
	user.Lastname = r.FormValue("Lastname")
	user.Email = r.FormValue("Email")
	user.Age, err = strconv.ParseUint(r.FormValue("Age"), 10, 32)

	// check data
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

	// insert new data into db
	db.QueryRow("INSERT INTO Users (firstname, lastname, email, age, created) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Firstname, user.Lastname, user.Email, user.Age, time.Now()).Scan(&user.ID)

	msg = fmt.Sprintf("The user %v was added", user.ID)
	fmt.Fprint(w, msg)
	log.Print(msg)
}

// Show user by ID
// http://127.0.0.1:8000/users/a5657a25-b62d-45f8-96f6-41aab04f9ec0
func showUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["key"]

	// check id
	if isUUID(id) != true {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		log.Printf("Incorrect ID %s", id)
		return
	}

	// select table row
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	user := tUser{}

	// scan data
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname,
		&user.Email, &user.Age, &user.Created)

	// handle error
	switch err {
	case sql.ErrNoRows:
		http.Error(w, "Not found", http.StatusNotFound)
		log.Print(err, id)
		return
	case nil:
		log.Printf("Request completed successfully with ID %s", id)
	default:
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Print(err, id)
		return
	}

	msg := fmt.Sprintf("User data:\n%v, %s, %s, %d, %v\n", user.ID,
		user.Firstname, user.Lastname, user.Age, user.Created)
	fmt.Fprint(w, msg)
}

// Edit user data by ID
func editUser(w http.ResponseWriter, r *http.Request) {

	var (
		user    tUser
		err     error
		id, msg string
	)

	// read keys
	id = r.FormValue("ID")
	user.Firstname = r.FormValue("Firstname")
	user.Lastname = r.FormValue("Lastname")
	user.Email = r.FormValue("Email")
	user.Age, err = strconv.ParseUint(r.FormValue("Age"), 10, 32)

	// check data
	switch {
	case isUUID(id) != true:
		http.Error(w, "Incorrect users ID!", http.StatusBadRequest)
		return
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

	// update row data
	_, err = db.Exec("UPDATE users SET firstname = $1, lastname = $2, email = $3, age = $4, created = $5 WHERE id = $6",
		user.Firstname, user.Lastname, user.Email, user.Age, time.Now(), id)
	if err != nil {
		msg = fmt.Sprintf("Failed to update user %s\n", id)
		http.Error(w, msg, http.StatusBadRequest)
		log.Print(err, msg)
		return
	}

	msg = "Users data successfully updated"
	fmt.Fprintln(w, msg)
	log.Print(msg)
}
