package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"regexp"
	"strings"

	"github.com/rav1L/book-machine/modules/db"
	uuid "github.com/satori/go.uuid"
)

func validateUserCredentials(r *http.Request, user *db.User) (err error) {
	reg := regexp.MustCompile(`^[\w]{8,}$`)
	fmt.Println(user.Login)
	if !reg.MatchString(user.Login) {
		errorHandler(statusInvalidParameters, "Invalid login: minimum length: 8, only latin and digits", &err)
		return
	}
	reg = regexp.MustCompile(`^[\S]{8,}$`)
	if !reg.MatchString(user.Password) {
		errorHandler(statusInvalidParameters, "Invalid password: minimum length: 8, no spaces, minimum 1 digit and 1 letter", &err)
		return
	}
	isLetterPresent, _ := regexp.MatchString(`(?i)[A-ZА-ЯЁ]`, user.Password)
	isDigitPresent, _ := regexp.MatchString(`[\d]`, user.Password)
	if !isLetterPresent || !isDigitPresent {
		errorHandler(statusInvalidParameters, "Invalid password: minimum length: 8, no spaces, minimum 1 digit and 1 letter", &err)
		return
	}
	return
}

func doesPasswordMatch(password1 string, password2 string) bool {
	return password1 == password2
}

func getLogin(token string) (login string, err error) {
	if token == "" {
		errorHandler(statusNotAuthorized, "", &err)
		return
	}
	login, err = myDB.Users().GetLogin(token)
	if err != nil && err != errNoRows {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	if login == "" {
		errorHandler(statusNotAuthorized, "", &err)
	}
	return
}

func getID(token string) (id string, err error) {
	if token == "" {
		errorHandler(statusNotAuthorized, "", &err)
		return
	}
	id, err = myDB.Users().GetID(token)
	if err != nil && err != errNoRows {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	if id == "" {
		errorHandler(statusNotAuthorized, "", &err)
	}
	return
}

func registerHandler(w http.ResponseWriter, r *http.Request) (err error) {
	switch r.Method {
	case "POST":
		err = r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		login := r.PostForm.Get(loginQuery)
		password := r.PostForm.Get(passwordQuery)
		user := &db.User{Login: login, Password: password}
		err = validateUserCredentials(r, user)
		if err != nil {
			return
		}
		token := r.PostForm.Get(tokenQuery)
		if token != config.AdminToken {
			user.AdminRights = false
		} else {
			user.AdminRights = true
		}
		var t string
		t, err = myDB.Users().Add(user)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.Contains(err.Error(), "unique"))
			if strings.Contains(err.Error(), "unique") {
				errorHandler(statusInvalidParameters, "user "+user.Login+" already exists", &err)
				return
			}
			errorHandler(statusNotExpected, "", &err)
			return
		}
		model := &outModel{}
		if user.AdminRights {
			model.Response = map[string]interface{}{tokenQuery: t, "message": "here's my man!"}
		} else {
			model.Response = map[string]interface{}{tokenQuery: t}
		}
		err = sendJSON(w, model)
		if err != nil {
			return
		}
	case "GET", "HEAD", "PUT", "PATCH", "DELETE", "OPTIONS", "TRACE", "CONNECT":
		errorHandler(statusUnimplementedMethod, "", &err)
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}

func authHandler(w http.ResponseWriter, r *http.Request) (err error) {
	switch r.Method {
	case "POST":
		err = r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		login := r.PostForm.Get(loginQuery)
		password := r.PostForm.Get(passwordQuery)
		user := &db.User{Login: login, Password: password}
		err = validateUserCredentials(r, user)
		if err != nil {
			return
		}
		var admin bool
		password, admin, err = myDB.Users().GetPassword(user.Login)
		if err != nil && err != errNoRows {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		if password == "" {
			errorHandler(statusNotAuthorized, "Invalid login", &err)
			return
		}
		if !doesPasswordMatch(user.Password, password) {
			errorHandler(statusNotAuthorized, "Wrong password", &err)
			return
		}
		var v4 uuid.UUID
		v4, err = uuid.NewV4()
		if err != nil {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		user.Token = v4.String()
		err = myDB.Users().UpdateToken(user.Login, user.Token)
		if err != nil {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		model := &outModel{}
		model.Response = map[string]interface{}{tokenQuery: user.Token, "admin": admin}
		err = sendJSON(w, model)
		if err != nil {
			return
		}
	case "GET", "HEAD", "PUT", "PATCH", "DELETE", "OPTIONS", "TRACE", "CONNECT":
		errorHandler(statusUnimplementedMethod, "", &err)
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}

func logoutHandler(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	token := path.Base(r.URL.Path)
	if token == "logout" {
		errorHandler(statusNotAuthorized, "", &err)
		return
	}
	switch r.Method {
	case "DELETE":
		err = myDB.Users().ClearToken(token)
		if err != nil {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		model := &outModel{}
		model.Response = map[string]interface{}{token: true}
		model.Response[token] = true
		err = sendJSON(w, model)
		if err != nil {
			return
		}
	case "GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS", "TRACE", "CONNECT":
		errorHandler(statusUnimplementedMethod, "", &err)
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}

func usersHandler(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()
	if err != nil {
		errorHandler(statusInvalidParameters, "", &err)
		return
	}
	switch r.Method {
	case "GET":
		token := r.Form.Get(tokenQuery)
		var admin bool
		if token != "" {
			admin, err = myDB.Users().IsAdmin(token)
			if err != nil {
				errorHandler(statusNotExpected, "", &err)
				return
			}
			if !admin {
				errorHandler(statusAccessDenied, "only admins", &err)
				return
			}
		} else {
			errorHandler(statusAccessDenied, "only admins", &err)
			return
		}
		model := &outModel{}
		var users []*db.User
		users, err = myDB.Users().GetList()
		if err != nil && err != errNoRows {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		model.Data = map[string]interface{}{"users": users}
		err = sendJSON(w, model)
	case "PUT", "POST":
		var user *db.User
		var id string
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		if r.Method == "POST" {
			id, err = myDB.Users().Add(user)
		} else {
			err = myDB.Users().Edit(user)
		}
		if err != nil && err != errNoRows {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		model := &outModel{}
		model.Data = map[string]interface{}{"id": id}
		err = sendJSON(w, model)
	case "OPTIONS":
		sendOptionsHeader(w, "OPTIONS, GET, POST, PUT")
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}
