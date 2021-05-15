package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"rest-api/internal/types"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// DB link
var DB *sql.DB

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
func checkKeys(r *http.Request) (types.User, error) {

	var (
		u   types.User
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

// AddUser writes new user data and returns UUID
func AddUser(w http.ResponseWriter, r *http.Request) {

	var (
		u   types.User
		err error
		msg string
	)

	u, err = checkKeys(r)
	if err != nil {
		msg = fmt.Sprintf("Error! %s", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// insert new data into DB
	DB.QueryRow("INSERT INTO Users (firstname, lastname, email, age, created) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		u.Firstname, u.Lastname, u.Email, u.Age, time.Now()).Scan(&u.ID)

	msg = fmt.Sprintf("The user %v was added", u.ID)
	fmt.Fprint(w, msg)
	log.Print(msg)
}

// DeleteUser erases user data by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var (
		err error
		msg string
	)

	// read key
	vars := mux.Vars(r)
	id := vars["key"]

	// check id
	if isUUID(id) != true {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		log.Printf("Incorrect ID %s", id)
		return
	}

	_, err = DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		msg = fmt.Sprintf("Failed to delete user %s\n", id)
		http.Error(w, msg, http.StatusBadRequest)
		log.Print(err, msg)
		return
	}

	msg = fmt.Sprintf("User with ID %s successfully deleted", id)
	log.Print(msg)
	fmt.Fprintln(w, msg)
}

// ShowAllUsers prints all users
func ShowAllUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := DB.Query("select * from users")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	users := []types.User{}

	for rows.Next() {
		u := types.User{}
		err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname,
			&u.Email, &u.Age, &u.Created)
		if err != nil {
			log.Print(err)
			continue
		}
		users = append(users, u)
	}

	msg := []string{}

	for _, u := range users {
		msg = append(msg, fmt.Sprintf("%v, %s, %s, %s, %d, %s\n",
			u.ID, u.Firstname, u.Lastname, u.Email, u.Age,
			u.Created.Format("2.1.2006 15:04")))
	}

	log.Print("All users printed")
	fmt.Fprint(w, "All users:\n"+strings.Join(msg, ""))
}

// ShowHomePage prints start page
func ShowHomePage(w http.ResponseWriter, r *http.Request) {

	var s string

	s = `<!DOCTYPE html>
<html lang="en">

    <head>
        <meta charset="UTF-8">
        <title>Hello!</title>
    </head>

    <body>
        <h1>Hello World!</h1>
        <p>This is a test home page</p>
    </body>

</html>`

	fmt.Fprint(w, s)
}

// ShowUser prints user data by ID
func ShowUser(w http.ResponseWriter, r *http.Request) {

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
	row := DB.QueryRow("SELECT * FROM users WHERE id = $1", id)

	u := types.User{}

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
		log.Printf("Printed user data with ID %s", id)
	default:
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Print(err, id)
		return
	}

	msg := fmt.Sprintf("User data:\n%v, %s, %s, %d, %s\n", u.ID,
		u.Firstname, u.Lastname, u.Age, u.Created.Format("2.1.2006 15:04"))
	fmt.Fprint(w, msg)
}

// EditUser writes new user data by ID
func EditUser(w http.ResponseWriter, r *http.Request) {

	var (
		u       types.User
		err     error
		id, msg string
	)

	// check id
	id = r.FormValue("id")
	if isUUID(id) != true {
		http.Error(w, "Error! Incorrect user ID!", http.StatusBadRequest)
	}

	// check keys
	u, err = checkKeys(r)
	if err != nil {
		msg = fmt.Sprintf("Error! %s", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// update row data
	_, err = DB.Exec("UPDATE users SET firstname = $1, lastname = $2, email = $3, age = $4, created = $5 WHERE id = $6",
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
