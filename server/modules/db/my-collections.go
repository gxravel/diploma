package db

import (
	"context"
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx"

	"github.com/jackc/pgx/pgxpool"
)

// MyCollections is
type MyCollections struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Collection string    `json:"collection"`
	BookID     string    `json:"book_id"`
	Added      time.Time `json:"added"`
}

// Collections is
type Collections struct {
	conn *pgxpool.Pool
}

// CollectionsMethods is
type CollectionsMethods interface {
	AddList([]*MyCollections) error
	Add(*MyCollections) (string, error)
	GetBookBelongs(string, string) ([]string, error)
	Get(string) ([]string, error)
	GetCollection(string, string, int, int) ([]*BookPreview, error)
	Delete(*MyCollections) error
	UpdateMethods
}

// AddList is
func (h *Collections) AddList(mc []*MyCollections) (err error) {
	for _, v := range mcList {
		_, err = h.Add(v)
		if err != nil {
			return
		}
	}
	return
}

// Add is
func (h *Collections) Add(mc *MyCollections) (id string, err error) {
	row := h.conn.QueryRow(context.Background(), `insert into my_collections (user_id, collection, book_id) 
	values ($1,$2,$3) returning id;`, mc.UserID, mc.Collection, mc.BookID)
	err = row.Scan(&id)
	return
}

// GetBookBelongs is
func (h *Collections) GetBookBelongs(user string, book string) (collections []string, err error) {
	rows, err := h.conn.Query(context.Background(), `select collection
	from my_collections
	where user_id=$1 and book_id=$2;`, user, book)
	if err != nil {
		return
	}
	defer rows.Close()
	var collection string
	for rows.Next() {
		err = rows.Scan(&collection)
		if err != nil {
			return
		}
		collections = append(collections, collection)
	}
	return
}

// Get is
func (h *Collections) Get(user string) (collections []string, err error) {
	collections = []string{"Буду читать", "Избранное"}
	rows, err := h.conn.Query(context.Background(), `select distinct collection
	from my_collections
	where user_id = $1
	order by collection;`, user)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		collection := ""
		err = rows.Scan(&collection)
		if err != nil {
			return
		}
		if collection != collections[0] && collection != collections[1] {
			collections = append(collections, collection)
		}
	}
	return
}

// GetCollection is
func (h *Collections) GetCollection(user string, collection string, onPage int, page int, column string) (books []*BookPreview, err error) {
	valid := regexp.MustCompile("^[.A-Za-z0-9_]+$")
	if !valid.MatchString(column) {
		return nil, errors.New("column is not valid")
	}
	rows, err := h.conn.Query(context.Background(), `select b.id, b.title, b.original_title, 
	b.publication, b.annotation, b.image, a.name
	from my_collections mc
	inner join books b on mc.book_id=b.id
	inner join authors a on b.author=a.id
	where mc.user_id=$1 and mc.collection=$2
	order by `+column+`
	limit $3 offset $4;`, user, collection, onPage, (page-1)*onPage)
	if err != nil {
		return
	}
	defer rows.Close()
	var rows2 pgx.Rows
	var publication time.Time
	for rows.Next() {
		name := ""
		b := &BookPreview{}
		err = rows.Scan(&b.ID, &b.Title, &b.OriginalTitle, &publication, &b.Annotation, &b.Image, &name)
		if err != nil {
			return
		}
		rows2, err = h.conn.Query(context.Background(), `select g.name 
			from book_genres bg
			inner join genres g on bg.genre_id=g.id
			where bg.book_id=$1;`, b.ID)
		if err != nil {
			return
		}
		defer rows2.Close()
		var genres []string
		for rows2.Next() {
			gName := ""
			err = rows2.Scan(&gName)
			if err != nil {
				return
			}
			genres = append(genres, gName)
		}
		b.Genres = strings.Join(genres, ", ")
		b.Author = name
		b.Publication = publication.Year()
		books = append(books, b)
	}
	return
}

// Delete is
func (h *Collections) Delete(mc *MyCollections) (err error) {
	_, err = h.conn.Exec(context.Background(), `delete from my_collections
	where user_id=$1 and collection=$2 and book_id=$3`, mc.UserID, mc.Collection, mc.BookID)
	return
}

// CreateTable is
func (h *Collections) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `
	create table my_collections (
		id uuid primary key default uuid_generate_v4(),
		user_id uuid NOT NULL REFERENCES users(id),
		collection varchar(100) NOT NULL default 'Новая коллекция',
		book_id uuid NOT NULL REFERENCES books(id),
		added TIMESTAMPTZ DEFAULT Now(),
		UNIQUE(user_id, collection, book_id)
		);`)
	return
}

// DropTable is
func (h *Collections) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists my_collections cascade;`)
	return
}

// Init is
func (h *Collections) Init() (err error) {
	log.Println("initializing my_collections.")
	err = h.AddList(mcList)
	return
}
