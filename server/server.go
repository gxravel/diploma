package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/rav1L/book-machine/modules/db"
)

const (
	statusOk                  = 200
	statusInvalidParameters   = 400
	statusNotAuthorized       = 401
	statusAccessDenied        = 403
	statusInvalidMethod       = 405
	statusNotExpected         = 500
	statusUnimplementedMethod = 501

	loginQuery      = "login"
	passwordQuery   = "password"
	tokenQuery      = "token"
	typeQuery       = "type"
	numberQuery     = "number"
	onPageQuery     = "onpage"
	pageQuery       = "page"
	maxMB           = 32 << 20
	dbDriver        = "postgres"
	host            = "localhost:8080"
	contentTypeJSON = "application/json; charset=utf-8"
	configName      = "config.json"
)

var (
	errNoRows    = pgx.ErrNoRows
	errCustomNil = errors.New("it will be ignored in the end but not before")
	clientError  *errorModel
	statusText   = map[int]string{
		statusInvalidParameters:   "Invalid parameters",
		statusNotAuthorized:       "Not authorized",
		statusAccessDenied:        "Access denied",
		statusInvalidMethod:       "Invalid request method",
		statusNotExpected:         "Not expected trouble",
		statusUnimplementedMethod: "The request method is not implemented",
		statusOk:                  ""}
	myDB   db.ISQL
	dbPath = os.Getenv("DATABASE_URL")
	config *configuration
)

type configuration struct {
	AdminToken string `json:"token"`
}

type outModel struct {
	Error    *errorModel            `json:"error,omitempty"`
	Response map[string]interface{} `json:"response,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

type errorModel struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func init() {
	myDB = &db.Handler{}
	err := myDB.Init(dbDriver, dbPath)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(configName)
	if err != nil {
		log.Fatal(err)
	}
	config = &configuration{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Fatal(err)
	}
	clientError = &errorModel{Code: 0}
}

func main() {
	http.HandleFunc("/register", makeHandler(registerHandler))
	http.HandleFunc("/auth", makeHandler(authHandler))
	http.HandleFunc("/logout/", makeHandler(logoutHandler))
	http.HandleFunc("/book/download/", makeHandler(downloadHandler))
	http.HandleFunc("/book/reviews/", makeHandler(reviewsHandler))
	http.HandleFunc("/book/rating/", makeHandler(ratingHandler))
	http.HandleFunc("/book/genres/", makeHandler(genresHandler))
	http.HandleFunc("/book/", makeHandler(booksHandler))
	http.HandleFunc("/book", makeHandler(booksHandler))
	http.HandleFunc("/my-collections/", makeHandler(myCollectionsIDHandler))
	http.HandleFunc("/my-collections", makeHandler(myCollectionsHandler))
	http.HandleFunc("/author/", makeHandler(authorsHandler))
	defer myDB.Disconnect()
	err := http.ListenAndServe(host, nil)
	log.Panic(err)
}

// errCustomNil is used for letting someHandler to know that an error was occured
// but it is not to be logged to the server

func makeHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil && err != errCustomNil {
			log.Printf("%+v", err)
		}
		if clientError.Code != 0 {
			if r.Method == "HEAD" {
				w.Header().Set("Content-Type", contentTypeJSON)
				w.WriteHeader(clientError.Code)
			} else {
				responseError(w)
			}
		}
		clientError.Code = 0
		clientError.Text = ""
	}
}

/* #region Auxiliary functions *********************************************************************************** */
func errorHandler(code int, text string, err *error) {
	var ok bool
	clientError.Text, ok = statusText[code]
	if !ok {
		errorHandler(statusNotExpected, "", err)
		return
	}
	clientError.Code = code
	if text != "" {
		clientError.Text += ": " + text
	}
	if code == statusNotExpected {
		*err = errors.WithStack(*err)
	} else {
		*err = errCustomNil
	}
}

func responseError(w http.ResponseWriter) {
	model := &outModel{}
	model.Error = clientError
	err := sendJSON(w, model)
	if err != nil {
		log.Println("wow!")
		http.Error(w, clientError.Text, clientError.Code)
	}
}

func sendJSON(w http.ResponseWriter, model *outModel) (err error) {
	modelJSON, err := json.Marshal(model)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	w.Header().Set("Content-Type", contentTypeJSON)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	_, err = w.Write(modelJSON)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	return
}

func sendOptionsHeader(w http.ResponseWriter, methods string) {
	w.Header().Set("Content-Type", contentTypeJSON)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", methods)
	w.WriteHeader(statusOk)
}

func sendFile(w http.ResponseWriter, dataPath string, fileName string) (err error) {
	var f *os.File
	f, err = os.Open(dataPath)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	var fi os.FileInfo
	fi, err = f.Stat()
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	log.Println("sending...")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Disposition", "attachment; filename="+url.PathEscape(fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("ContentLength", fmt.Sprint(fi.Size()))
	_, err = io.Copy(w, f)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
	}
	return
}

func readMultipartFile(r *http.Request, fpath string, name string) (err error) {
	var file multipart.File
	var handler *multipart.FileHeader
	file, handler, err = r.FormFile("file")
	if err != nil {
		if strings.Contains(err.Error(), "http: no such file") {
			return nil
		}
		errorHandler(statusNotExpected, "", &err)
		return
	}
	defer file.Close()
	path := filepath.Join(fpath, name) + filepath.Ext(handler.Filename)
	var f *os.File
	f, err = os.Create(path)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	return
}

func readMulitpart(r *http.Request) (book *db.Book, token string, err error) {
	err = r.ParseMultipartForm(maxMB)
	if err != nil {
		errorHandler(statusInvalidParameters, "Memory limit size was overloaded", &err)
		return
	}
	bookString := r.Form.Get("book")
	token = r.Form.Get("token")
	var admin bool
	admin, err = myDB.Users().IsAdmin(token)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	if !admin {
		errorHandler(statusAccessDenied, "only admins", &err)
		return
	}
	err = json.NewDecoder(strings.NewReader(bookString)).Decode(&book)
	if err != nil {
		errorHandler(statusNotExpected, "", &err)
		return
	}
	var name = book.ID
	err = readMultipartFile(r, "./data", name)
	return
}

/* #endregion *************************************************************************************************** */
