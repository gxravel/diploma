package main

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/rav1L/book-machine/modules/db"
)

func reviewsHandler(w http.ResponseWriter, r *http.Request) (err error) {
	model := &outModel{}
	urlBase := path.Base(r.URL.Path)
	log.Println("reviews handler: ", r.URL.Path)
	switch r.Method {
	case "GET":
		r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		if urlBase != "reviews" {
			var reviews []*db.Review
			reviews, err = myDB.Reviews().Get(urlBase)
			if err != nil && err != errNoRows {
				errorHandler(statusNotExpected, "", &err)
				return
			}
			model.Data = map[string]interface{}{"reviews": reviews}
			err = sendJSON(w, model)
		} else {
			errorHandler(statusInvalidParameters, "book id in url is required", &err)
		}
	case "POST":
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
			header := r.PostForm.Get("header")
			reviewText := r.PostForm.Get("review_text")
			if reviewText == "" {
				errorHandler(statusInvalidParameters, "review text is required", &err)
			}
			review := &db.Review{BookID: urlBase, UserID: uid, Header: header,
				ReviewText: reviewText, DateAdded: time.Now()}
			err = myDB.Reviews().Add(review)
			if err != nil && err != errNoRows {
				errorHandler(statusNotExpected, "", &err)
				return
			}
			err = sendJSON(w, model)
		} else {
			errorHandler(statusAccessDenied, "only authorized users", &err)
		}
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}
