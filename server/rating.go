package main

import (
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/rav1L/book-machine/modules/db"
)

func ratingHandler(w http.ResponseWriter, r *http.Request) (err error) {
	model := &outModel{}
	urlBase := path.Base(r.URL.Path)
	log.Println("rating handler: ", r.URL.Path)
	switch r.Method {
	case "GET":
		r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		if urlBase != "rating" {
			var userValue float64 = 0
			var avgValue float64
			var number int
			var userID string
			token := r.Form.Get("token")
			if token != "" {
				userID, err = getID(token)
				if err != nil {
					return
				}
				userValue, err = myDB.Rating().GetPersonal(urlBase, userID)
				if err != nil && err != errNoRows {
					errorHandler(statusNotExpected, "", &err)
					return
				}
			}
			avgValue, number, err = myDB.Rating().GetAvgValue(urlBase)
			if err != nil && err != errNoRows {
				errorHandler(statusNotExpected, "", &err)
				return
			}
			rating := map[string]interface{}{"user_value": userValue, "avg_value": avgValue, "number": number}
			model.Data = map[string]interface{}{"rating": rating}
			err = sendJSON(w, model)
		} else {
			errorHandler(statusInvalidParameters, "book id in url is required", &err)
		}
	case "POST", "PUT":
		r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		token := r.PostForm.Get(tokenQuery)
		if token != "" {
			var uid string
			uid, err = getID(token)
			if err != nil {
				return
			}
			var value int
			value, err = strconv.Atoi(r.PostForm.Get("user_value"))
			if err != nil {
				errorHandler(statusInvalidParameters, "user value must be from 1 to 5", &err)
				return
			}
			rating := &db.Rating{BookID: urlBase, UserID: uid, Value: value, DateAdded: time.Now()}
			if r.Method == "POST" {
				err = myDB.Rating().Add(rating)
			} else {
				err = myDB.Rating().Edit(rating)
			}
			if err != nil && err != errNoRows {
				errorHandler(statusNotExpected, "", &err)
				return
			}
			err = sendJSON(w, model)
		} else {
			errorHandler(statusAccessDenied, "only authorized users", &err)
		}
	case "OPTIONS":
		sendOptionsHeader(w, "OPTIONS, GET, POST, PUT")
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}
