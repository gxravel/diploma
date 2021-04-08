package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/rav1L/book-machine/modules/db"
	dc "github.com/rav1L/book-machine/modules/draw-cover"
)

const (
	defaultExt = ".fb2"
	dataPath   = "data/"
)

func booksHandler(w http.ResponseWriter, r *http.Request) (err error) {
	switch r.Method {
	case "GET":
		r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		search := r.Form.Get("search")
		fmt.Println(search)
		model := &outModel{}
		urlBase := path.Base(r.URL.Path)
		log.Println(urlBase)
		if urlBase == "book" {
			if search != "" {
				var books []*db.BookPreview
				books, err = myDB.Books().Search(search)
				if err != nil && err != errNoRows {
					errorHandler(statusNotExpected, "", &err)
					return
				}
				model.Data = map[string]interface{}{"books": books}
			} else {
				model.Data = map[string]interface{}{"books": ""}
			}
		} else {
			books := make([]*db.Book, 1)
			books[0], err = myDB.Books().Get(urlBase)
			if err != nil && err != errNoRows {
				errorHandler(statusNotExpected, "", &err)
				return
			}
			dc.Draw(books[0].Title, books[0].ID)
			token := r.Form.Get(tokenQuery)
			if token != "" {
				var uid string
				uid, err = getID(token)
				if err != nil {
					return
				}
				var collections []string
				collections, err = myDB.Collections().GetBookBelongs(uid, urlBase)
				if err != nil && err != errNoRows {
					errorHandler(statusNotExpected, "", &err)
					return
				}
				belongs := map[string]bool{"Буду читать": false, "Избранное": false}
				for _, v := range collections {
					belongs[v] = true
				}
				log.Println("booksHandler: ", collections, belongs)
				model.Data = map[string]interface{}{"books": books, "belongs": belongs}
			} else {
				model.Data = map[string]interface{}{"books": books}
			}
		}
		err = sendJSON(w, model)
		if err != nil {
			return
		}
	case "PUT", "POST":
		var book *db.Book
		var id string
		book, _, err = readMulitpart(r)
		if err != nil {
			return
		}
		if r.Method == "POST" {
			id, err = myDB.Books().Add(book)
		} else {
			err = myDB.Books().Edit(book)
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

func downloadHandler(w http.ResponseWriter, r *http.Request) (err error) {

	id := path.Base(r.URL.Path)
	if id == "download" {

		errorHandler(statusInvalidParameters, "id is missing or it is `download` - offensive and inappropriate value", &err)
		return
	}
	log.Println("r.Method", r.Method)

	switch r.Method {
	case "GET":

		err = r.ParseForm()
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		ext := r.Form.Get("ext")
		if ext == "" {
			ext = defaultExt
		}
		var filename string
		filename, err = myDB.Books().GetFileName(id)
		if err != nil {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		log.Println(filename)
		err = sendFile(w, filepath.Join(dataPath, id)+ext, filename)
	}
	return
}
