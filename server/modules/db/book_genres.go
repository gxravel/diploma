package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/pgxpool"
)

// BookGenres is
type BookGenres struct {
	BookID  string `json:"book_id"`
	GenreID string `json:"genre_id"`
}

// BookGenresHandler is
type BookGenresHandler struct {
	conn *pgxpool.Pool
}

// BookGenresMethods is
type BookGenresMethods interface {
	AddList([]*BookGenres) error
	Add(*BookGenres) error
	UpdateMethods
}

// AddList is
func (h *BookGenresHandler) AddList(bgList []*BookGenres) (err error) {
	for _, v := range bgList {
		err = h.Add(v)
		if err != nil {
			return
		}
	}
	return
}

// Add is
func (h *BookGenresHandler) Add(bg *BookGenres) (err error) {
	_, err = h.conn.Exec(context.Background(), `insert into book_genres (book_id, genre_id) 
	values ($1,$2);`, bg.BookID, bg.GenreID)
	if err != nil {
		return
	}
	return
}

// CreateTable is
func (h *BookGenresHandler) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `
	create table book_genres (
		id uuid primary key default uuid_generate_v4(),
		book_id uuid NOT NULL REFERENCES books(id),
		genre_id uuid NOT NULL REFERENCES genres(id)
		);`)
	return
}

// DropTable is
func (h *BookGenresHandler) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists book_genres cascade;`)
	return
}

// Init is
func (h *BookGenresHandler) Init() (err error) {
	log.Println("initializing book_genres.")
	err = h.AddList(bgList)
	return
}
