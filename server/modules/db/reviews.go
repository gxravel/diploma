package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/pgxpool"
)

// Review is
type Review struct {
	ID         string    `json:"id"`
	BookID     string    `json:"book_id"`
	UserID     string    `json:"user_id"`
	Header     string    `json:"header"`
	ReviewText string    `json:"review_text"`
	DateAdded  time.Time `json:"date_added"`
	Username   string    `json:"username"`
}

// Reviews is
type Reviews struct {
	conn *pgxpool.Pool
}

// ReviewsMethods is
type ReviewsMethods interface {
	Add(*Review) error
	Get(string) ([]*Review, error)
	UpdateMethods
}

// Add is
func (h Reviews) Add(obj *Review) (err error) {
	_, err = h.conn.Exec(context.Background(), `insert into reviews (book_id, user_id, header, review_text, 
		date_added) values ($1,$2,$3,$4,$5)`, obj.BookID, obj.UserID, obj.Header, obj.ReviewText, obj.DateAdded)
	return
}

// CreateTable is
func (h Reviews) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `
	create table reviews (
		id uuid primary key default uuid_generate_v4(),
		book_id uuid NOT NULL REFERENCES books(id),
		user_id uuid NOT NULL REFERENCES users(id),
		header varchar(255),
		review_text text NOT NULL,
		date_added timestamp
		);`)
	return
}

// DropTable is
func (h Reviews) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists reviews cascade;`)
	return
}

// Get is
func (h Reviews) Get(id string) (objList []*Review, err error) {
	rows, err := h.conn.Query(context.Background(), `select u.login, header, review_text, date_added 
	from reviews r left join users u on r.user_id=u.id
	where book_id = $1`, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		obj := &Review{BookID: id}
		err = rows.Scan(&obj.Username, &obj.Header, &obj.ReviewText, &obj.DateAdded)
		if err != nil {
			return
		}
		objList = append(objList, obj)
	}
	return
}
