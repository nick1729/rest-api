package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// UUID validation
func isUUID(u string) bool {

	var err error

	_, err = uuid.Parse(u)
	if err != nil {
		return false
	}

	return true
}

// Name validation
func isName(n string) bool {

	var (
		re   *regexp.Regexp
		expr string
		ok   bool
	)

	expr = `^[A-Z][a-z]+${1,64}`
	re = regexp.MustCompile(expr)
	ok = re.MatchString(n)

	return ok
}

// Email validation
func isEmail(e string) bool {

	var (
		re   *regexp.Regexp
		expr string
		ok   bool
	)

	expr = `^((([0-9A-Za-z]{1}[-0-9A-z\.]{1,64}[0-9A-Za-z]{1}))@([-A-Za-z]{1,32}\.){1,2}[-A-Za-z]{2,4})$`
	re = regexp.MustCompile(expr)
	ok = re.MatchString(e)

	return ok
}

// Checking keys
func checkKeys(r *http.Request) (tUser, error) {

	var (
		u   tUser
		err error
	)

	// read keys
	u.Firstname = r.FormValue("firstname")
	u.Lastname = r.FormValue("lastname")
	u.Email = r.FormValue("email")
	u.Age, err = strconv.ParseUint(r.FormValue("age"), 10, 32)

	// check data
	switch {
	case isName(u.Firstname) != true:
		err = errors.New("Incorrect user firstname!")
		return u, err
	case isName(u.Lastname) != true:
		err = errors.New("Incorrect user lastname!")
		return u, err
	case isEmail(u.Email) != true:
		err = errors.New("Incorrect user email!")
		return u, err
	case err != nil || u.Age < 1 || u.Age > 160:
		err = errors.New("Incorrect user age!")
		return u, err
	}

	return u, nil
}

// Add new user
// /users?firstname=Gena&lastname=Ivanov&email=qweqw@mail.gg&age=29
func addUser(w http.ResponseWriter, r *http.Request) {

	var (
		u   tUser
		err error
		msg string
	)

	u, err = checkKeys(r)
	if err != nil {
		msg = fmt.Sprintf("Error! %s", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// insert new data into db
	db.QueryRow("INSERT INTO Users (firstname, lastname, email, age, created) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		u.Firstname, u.Lastname, u.Email, u.Age, time.Now()).Scan(&u.ID)

	msg = fmt.Sprintf("The user %v was added", u.ID)
	fmt.Fprint(w, msg)
	log.Print(msg)
}

// Show user by ID
// /users/a5657a25-b62d-45f8-96f6-41aab04f9ec0
func showUser(w http.ResponseWriter, r *http.Request) {

	// read key
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

	u := tUser{}

	// scan data
	err := row.Scan(&u.ID, &u.Firstname, &u.Lastname,
		&u.Email, &u.Age, &u.Created)

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

	msg := fmt.Sprintf("User data:\n%v, %s, %s, %d, %v\n", u.ID,
		u.Firstname, u.Lastname, u.Age, u.Created)
	fmt.Fprint(w, msg)
}

// Edit user data by ID
// /users/?id=a5657a25-b62d-45f8-96f6-41aab04f9ec0&firstname=Qwe&lastname=Rty&email=qwe@rty.gg&age=23
func editUser(w http.ResponseWriter, r *http.Request) {

	var (
		u       tUser
		err     error
		id, msg string
	)

	id = r.FormValue("id")
	if isUUID(id) != true {
		http.Error(w, "Error! Incorrect user ID!", http.StatusBadRequest)
	}

	u, err = checkKeys(r)
	if err != nil {
		msg = fmt.Sprintf("Error! %s", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// update row data
	_, err = db.Exec("UPDATE users SET firstname = $1, lastname = $2, email = $3, age = $4, created = $5 WHERE id = $6",
		u.Firstname, u.Lastname, u.Email, u.Age, time.Now(), id)
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
