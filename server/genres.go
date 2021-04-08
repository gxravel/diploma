package main

import (
	"log"
	"net/http"

	"github.com/rav1L/book-machine/modules/db"
)

func genresHandler(w http.ResponseWriter, r *http.Request) (err error) {
	model := &outModel{}
	log.Println("genres handler: ")
	switch r.Method {
	case "GET":
		if err != nil {
			errorHandler(statusInvalidParameters, "", &err)
			return
		}
		var genres []*db.Genre
		genres, err = myDB.Genres().GetList()
		if err != nil && err != errNoRows {
			errorHandler(statusNotExpected, "", &err)
			return
		}
		model.Data = map[string]interface{}{"genres": genres}
		err = sendJSON(w, model)
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}
