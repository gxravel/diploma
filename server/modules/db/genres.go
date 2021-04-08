package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/pgxpool"
)

// Genre is
type Genre struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

// Genres is
type Genres struct {
	conn *pgxpool.Pool
}

// GenreMethods is
type GenreMethods interface {
	AddList([]*Genre) error
	Add(*Genre) error
	GetList() ([]*Genre, error)
	UpdateMethods
}

// AddList is
func (h *Genres) AddList(gList []*Genre) (err error) {
	for _, v := range gList {
		err = h.Add(v)
		if err != nil {
			return
		}
	}
	return
}

// Add is
func (h *Genres) Add(g *Genre) (err error) {
	if g.ID != "" {
		_, err = h.conn.Exec(context.Background(), `insert into genres (id, name, parent) 
		values ($1,$2,NULLIF($3,'')::uuid);`, g.ID, g.Name, g.Parent)
	} else {
		_, err = h.conn.Exec(context.Background(), `insert into genres (name, parent) 
		values ($1,NULLIF($2,'')::uuid);`, g.Name, g.Parent)
	}
	return
}

// GetList is
func (h *Genres) GetList() (list []*Genre, err error) {
	var nullString sql.NullString
	rows, err := h.conn.Query(context.Background(), `select id, name, parent 
	from genres`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		g := &Genre{}
		err = rows.Scan(&g.ID, &g.Name, &nullString)
		if err != nil {
			return
		}
		if nullString.Valid {
			g.Parent = nullString.String
		} else {
			g.Parent = ""
		}
		list = append(list, g)
	}
	return
}

// CreateTable is
func (h *Genres) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `
	create table genres (
		id uuid primary key default uuid_generate_v4(),
		name varchar(40) NOT NULL default 'unspecified',
		parent uuid default null REFERENCES genres(id) check (parent != id)
		);`)
	return
}

// DropTable is
func (h *Genres) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists genres;`)
	return
}

// Init is
func (h *Genres) Init() (err error) {
	log.Println("initializing genres.")
	err = h.AddList(gList)
	return
}
