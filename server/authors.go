package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/rav1L/book-machine/modules/db"
)

func authorsHandler(w http.ResponseWriter, r *http.Request) (err error) {
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
		if urlBase == "author" {
			if search != "" {
				var authors []*db.Author
				authors, err = myDB.Authors().Search(search)
				if err != nil && err != errNoRows {
					errorHandler(statusNotExpected, "", &err)
					return
				}
				model.Data = map[string]interface{}{"authors": authors}
			} else {
				model.Data = map[string]interface{}{"authors": ""}
			}
		} 
		err = sendJSON(w, model)
		if err != nil {
			return
		}
	default:
		errorHandler(statusInvalidMethod, "", &err)
	}
	return
}
