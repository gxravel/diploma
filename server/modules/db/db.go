package db

import (
	"context"

	"github.com/jackc/pgx/pgxpool"
)

// Filter is the parameters for building queries
type Filter struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}

// ISQL is the interface of sql database primarily for flexibility and mocking
type ISQL interface {
	Connect() error
	Disconnect()
	Init(string, string) error
	Authors() *Authors
	Users() *Users
	Books() *Books
	Collections() *Collections
	Genres() *Genres
	BookGenres() *BookGenresHandler
	Reviews() *Reviews
	Rating() *Ratings
}

// UpdateMethods is
type UpdateMethods interface {
	CreateTable() error
	DropTable() error
	Init() error
}

// Handler is sql database tool to work with sqlDriver
type Handler struct {
	conn   *pgxpool.Pool
	path   string
	driver string
}

// Connect creates connection to the database
func (h *Handler) Connect() (err error) {
	h.conn, err = pgxpool.Connect(context.Background(), h.path)
	return
}

// Disconnect closes connection of the database
func (h *Handler) Disconnect() {
	h.conn.Close()
}

// Init creates connection to the database and prepares the statements
func (h *Handler) Init(driver string, path string) (err error) {
	h.driver = driver
	h.path = path
	err = h.Connect()
	if err != nil {
		return
	}
	return
}

// Authors is
func (h *Handler) Authors() *Authors {
	return &Authors{h.conn}
}

// Books is
func (h *Handler) Books() *Books {
	return &Books{h.conn}
}

// Collections is
func (h *Handler) Collections() *Collections {
	return &Collections{h.conn}
}

// Users is
func (h *Handler) Users() *Users {
	return &Users{h.conn}
}

// BookGenres is
func (h *Handler) BookGenres() *BookGenresHandler {
	return &BookGenresHandler{h.conn}
}

// Genres is
func (h *Handler) Genres() *Genres {
	return &Genres{h.conn}
}

// Reviews is
func (h *Handler) Reviews() *Reviews {
	return &Reviews{h.conn}
}

// Rating is
func (h *Handler) Rating() *Ratings {
	return &Ratings{h.conn}
}
