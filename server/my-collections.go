package main

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/rav1L/book-machine/modules/db"
)

const (
	defaultCollection = "Буду читать"
)

var (
	defaultCollections = []string{"Буду читать", "Избранное"}
)

func myCollectionsHandler(w http.ResponseWriter, r *http.Request) (err error) {
	switch r.Method {
	case "GET":
		r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		token := r.Form.Get("token")
		var (
			user        string
			collections []string
		)
		user, err = getID(token)
		if err != nil {
			return
		}
		collections, err = myDB.Collections().Get(user)
		if err != nil && err != errNoRows {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		model := &outModel{}
		model.Data = map[string]interface{}{"collections": collections}
		err = sendJSON(w, model)
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}

func myCollectionsIDHandler(w http.ResponseWriter, r *http.Request) (err error) {
	urlBase := path.Base(r.URL.Path)
	if urlBase == "my-collections" {
		errorHandler(statusInvalidParameters, "", &err)
		return
	}
	switch r.Method {
	case "GET", "POST", "DELETE":
		r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		model := &outModel{}
		var (
			token      string
			collection string
			user       string
			books      []*db.BookPreview
		)
		if r.Method == "POST" {
			token = r.PostForm.Get(tokenQuery)
			collection = r.PostForm.Get("collection")
		} else {
			token = r.Form.Get(tokenQuery)
			collection = r.Form.Get("collection")

		}
		log.Println("myCollectionsIDHandler: ", token, urlBase)
		user, err = getID(token)
		if err != nil {
			return
		}
		log.Println("myCollectionsIDHandler: ", user)

		switch r.Method {
		case "GET":
			var (
				onPage int
				page   int
			)
			onPage, err = strconv.Atoi(r.Form.Get(onPageQuery))
			if err != nil {
				onPage = 10
				err = nil
			}
			page, err = strconv.Atoi(r.Form.Get(pageQuery))
			if err != nil {
				page = 1
				err = nil
			}
			column := r.Form.Get("column")
			if column == "" {
				column = "mc.title"
			}
			books, err = myDB.Collections().GetCollection(user, urlBase, onPage, page, column)
			if err != nil && err != errNoRows {
				errorHandler(statusNotExpected, "", &err)
				return
			}
			model := &outModel{}
			model.Data = map[string]interface{}{"books": books}
			err = sendJSON(w, model)
		case "POST", "DELETE":
			if collection == "" {
				errorHandler(statusInvalidParameters, "a collection must be provided", &err)
				return
			}
			mc := &db.MyCollections{UserID: user, BookID: urlBase, Collection: collection}
			if r.Method == "POST" {
				_, err = myDB.Collections().Add(mc)
				if err != nil {
					errorHandler(statusNotExpected, "", &err)
					return
				}
			} else {
				err = myDB.Collections().Delete(mc)
				if err != nil {
					errorHandler(statusNotExpected, "", &err)
					return
				}
			}
			model.Response = map[string]interface{}{"message": "OK"}
			err = sendJSON(w, model)
		}
	case "OPTIONS":
		sendOptionsHeader(w, "OPTIONS, GET, POST, DELETE")
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}
