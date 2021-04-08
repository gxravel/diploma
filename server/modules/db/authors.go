package db

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/pgxpool"
)

// Author is
type Author struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	OriginalName     string    `json:"original_name"`
	BirthDate        time.Time `json:"birth_date"`
	DeathDate        time.Time `json:"death_date"`
	UnknownDeathDate bool      `json:"unknown_dd"`
}

// Authors is
type Authors struct {
	conn *pgxpool.Pool
}

// AuthorsMethods is
type AuthorsMethods interface {
	AddList([]*Author) error
	Add(*Author) (string, error)
	GetName(string) (string, error)
	Get(string) (*Author, error)
	Search(string) ([]*Author, error)
	UpdateMethods
}

// AddList is
func (h Authors) AddList(aList []*Author) (err error) {
	for _, v := range aList {
		_, err = h.Add(v)
		if err != nil {
			return
		}
	}
	return
}

// Add is
func (h Authors) Add(a *Author) (id string, err error) {
	result, err := h.conn.Exec(context.Background(), `insert into authors (name, original_name, 
		birth_date, death_date) values ($1,$2,$3,$4) returning id;`, a.Name, a.OriginalName, a.BirthDate, a.DeathDate)
	if err != nil {
		return
	}
	id = result.String()
	return
}

// CreateTable is
func (h Authors) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `
	create table authors (
		id uuid primary key default uuid_generate_v4(),
		name varchar(100) NOT NULL,
		original_name varchar(100),
		birth_date DATE,
		death_date DATE
		);`)
	return
}

// DropTable is
func (h Authors) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists authors cascade;`)
	return
}

// GetName is
func (h Authors) GetName(id string) (name string, err error) {
	row := h.conn.QueryRow(context.Background(), `select name from authors where id=$1;`, id)
	err = row.Scan(&name)
	return
}

// Get is
func (h Authors) Get(id string) (a *Author, err error) {
	row := h.conn.QueryRow(context.Background(), `select name, birth_date, death_date 
	from authors where id = $1`, id)
	a = &Author{ID: id}
	err = row.Scan(&a.Name, &a.BirthDate, &a.DeathDate)
	return
}

// Init is
func (h Authors) Init() (err error) {
	log.Println("initializing authors.")
	err = h.AddList(aList)
	return
}

// Search is
func (h Authors) Search(value string) (aList []*Author, err error) {
	value = "%" + strings.ToLower(value) + "%"
	rows, err := h.conn.Query(context.Background(), `select a.id, a.name, a.original_name
	from authors a 
	where lower(a.name) LIKE $1 OR lower(a.original_name) LIKE $1
	order by a.name  
	limit $2`, value, searchNumber)
	if err != nil {
		return
	}
	defer rows.Close()
	var a *Author
	for rows.Next() {
		a = &Author{}

		err = rows.Scan(&a.ID, &a.Name, &a.OriginalName)
		if err != nil {
			return
		}
		aList = append(aList, a)
	}
	return
}
