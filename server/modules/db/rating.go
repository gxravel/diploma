package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/pgxpool"
)

// Rating is
type Rating struct {
	ID        string    `json:"id"`
	BookID    string    `json:"book_id"`
	UserID    string    `json:"user_id"`
	Value     int       `json:"value"`
	DateAdded time.Time `json:"date_added"`
}

// Ratings is
type Ratings struct {
	conn *pgxpool.Pool
}

// RatingsMethods is
type RatingsMethods interface {
	Add(*Rating) error
	Edit(*Rating) error
	GetAvgValue(string) (float64, int, error)
	GetPersonal(string, string) (float64, error)
	UpdateMethods
}

// Add is
func (h Ratings) Add(obj *Rating) (err error) {
	_, err = h.conn.Exec(context.Background(), `insert into rating (book_id, user_id, value, 
		date_added) values ($1,$2,$3,$4)`, obj.BookID, obj.UserID, obj.Value, obj.DateAdded)
	return
}

// Edit is
func (h Ratings) Edit(obj *Rating) (err error) {
	_, err = h.conn.Exec(context.Background(), `update rating set book_id=$1, value=$2, date_added=$3
	where user_id=$4`, obj.BookID, obj.Value, obj.DateAdded, obj.UserID)
	return
}

// CreateTable is
func (h Ratings) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `
	create table rating (
		id uuid primary key default uuid_generate_v4(),
		book_id uuid NOT NULL REFERENCES books(id),
		user_id uuid NOT NULL REFERENCES users(id),
		value int NOT NULL,
		date_added timestamp,
		UNIQUE(book_id, user_id)
		);`)
	return
}

// DropTable is
func (h Ratings) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists rating cascade;`)
	return
}

// GetAvgValue is
func (h Ratings) GetAvgValue(bookID string) (value float64, number int, err error) {
	var tValue sql.NullFloat64
	// row := h.conn.QueryRow(context.Background(), `select * from
	// (select avg(value)
	// from rating
	// where book_id = $1) avg_value,
	// (SELECT reltuples::bigint AS estimate
	// 	FROM   pg_class
	// 	WHERE  oid = 'public.rating'::regclass) number;`, bookID)
	row := h.conn.QueryRow(context.Background(), `
	select avg(value), COUNT(*)
	from rating
	where book_id = $1`, bookID)
	err = row.Scan(&tValue, &number)
	if err != nil {
		return
	}
	if tValue.Valid {
		value = tValue.Float64
	} else {
		value = 0
	}
	return
}

// GetPersonal is
func (h Ratings) GetPersonal(bookID string, userID string) (rating float64, err error) {
	var tValue sql.NullFloat64
	row := h.conn.QueryRow(context.Background(), `select value 
	from rating
	where book_id = $1 and user_id = $2`, bookID, userID)
	err = row.Scan(&tValue)
	if err != nil {
		return
	}
	if tValue.Valid {
		rating = tValue.Float64
	} else {
		rating = 0
	}
	return
}
