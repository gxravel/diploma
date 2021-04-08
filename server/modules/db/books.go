package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/pgxpool"
)

const (
	searchNumber = 10
)

// Book is
type Book struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	OriginalTitle    string    `json:"original_title"`
	Author           *Author   `json:"author"`
	Genres           []*Genre  `json:"genres"`
	OriginalLanguage string    `json:"original_language"`
	WritingStart     time.Time `json:"writing_start"`
	WritingEnd       time.Time `json:"writing_end"`
	Publication      time.Time `json:"publication"`
	Rating           []*Rating `json:"rating"`
	Annotation       string    `json:"annotation"`

	Source string `json:"source"`
	Image  string `json:"image"`
}

// Books is
type Books struct {
	conn *pgxpool.Pool
}

// BookPreview is
type BookPreview struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	Author        string  `json:"author"`
	Publication   int     `json:"publication"`
	Annotation    string  `json:"annotation"`
	RatingValue   float64 `json:"rating_value"`
	Genres        string  `json:"genres"`

	Image string `json:"image"`
}

// BooksMethods is
type BooksMethods interface {
	AddList([]*Book) error
	Add(*Book) (string, error)
	Edit(*Book) error
	GetList(int, int) ([]*Book, error)
	GetOrderedList(int, int, string, bool) ([]*Book, error)
	GetFileName(string) (string, error)
	Get(string) (*Book, error)
	Remove(string) error
	Search(string) ([]*BookPreview, error)
	UpdateMethods
}

// AddList is
func (h Books) AddList(bList []*Book) (err error) {
	for _, v := range bList {
		_, err = h.Add(v)
		if err != nil {
			return
		}
	}
	return
}

// Add is
func (h Books) Add(b *Book) (id string, err error) {
	row := h.conn.QueryRow(context.Background(), `insert into books (title, original_title, author, 
		original_language, writing_start, writing_end, publication, annotation, source, image) 
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id;`,
		b.Title, b.OriginalTitle, b.Author.ID, b.OriginalLanguage, b.WritingStart,
		b.WritingEnd, b.Publication, b.Annotation, b.Source, b.Image)
	err = row.Scan(&id)
	if err != nil {
		return
	}
	for _, v := range b.Genres {
		_, err = h.conn.Exec(context.Background(), `insert into book_genres (book_id, genre_id)
		values ($1, $2);`, id, v.ID)
		if err != nil {
			return
		}
	}
	return
}

// CreateTable is
func (h Books) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `
	create table books (
		id uuid primary key default uuid_generate_v4(),
		title varchar(50) NOT NULL,
		original_title varchar(50),
		author uuid NOT NULL REFERENCES authors(id),
		original_language varchar(20),
		writing_start date,
		writing_end date,
		publication date,
		annotation text,
		source varchar(100),
		image varchar(100),
		deleted bool default 'f'
		);`)
	return
}

// DropTable is
func (h Books) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists books cascade;`)
	return
}

// Edit is
func (h Books) Edit(b *Book) (err error) {
	_, err = h.conn.Exec(context.Background(), `update books set title=$1, original_title=$2, author=$3, 
	original_language=$4, writing_start=$5, writing_end=$6, publication=$7, 
	annotation=$8, source=$9, image=$10
	where id=$11;`, b.Title, b.OriginalTitle, b.Author.ID, b.OriginalLanguage, b.WritingStart,
		b.WritingEnd, b.Publication, b.Annotation, b.Source, b.Image, b.ID)
	if err != nil {
		return
	}
	_, err = h.conn.Exec(context.Background(), `delete from book_genres where book_id=$1;`, b.ID)
	if err != nil {
		return
	}
	for _, v := range b.Genres {
		_, err = h.conn.Exec(context.Background(), `insert into book_genres (book_id, genre_id) 
			values ($1, $2);`, b.ID, v.ID)
		if err != nil {
			return
		}
	}
	return
}

// // GetList is
// func (h Books) GetList(onPage int, page int) (bList []*Book, err error) {
// 	return h.GetOrderedList(onPage, page, "title", true)
// }

// // GetOrderedList is
// func (h Books) GetOrderedList(onPage int, page int, column string, asc bool) (bList []*Book, err error) {
// 	var sortingOrder string
// 	if asc {
// 		sortingOrder = "ASC"
// 	} else {
// 		sortingOrder = "DESC"
// 	}
// 	column = "b." + column
// 	rows, err := h.conn.Query(context.Background(), `select b.id, b.title, b.author, b.year,
// 	b.genre, b.image, a.name
// 	from books b
// 	left join authors a on b.author=a.id
// 	where NOT b.deleted
// 	order by $3`+sortingOrder+`
// 	limit $1 offset $2;`, onPage, (page-1)*onPage, column)
// 	if err != nil {
// 		return
// 	}
// 	defer rows.Close()
// 	bList = make([]*Book, onPage)
// 	var b *Book
// 	for rows.Next() {
// 		b = &Book{}
// 		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year, &b.Genre, &b.Image, &b.AuthorName)
// 		if err != nil {
// 			return
// 		}
// 		bList = append(bList, b)
// 	}
// 	return
// }

// GetFileName is
func (h Books) GetFileName(id string) (filename string, err error) {
	row := h.conn.QueryRow(context.Background(), `select b.title, a.name
	from books b 
	left join authors a on b.author=a.id
	where NOT b.deleted and b.id = $1`, id)
	var title string
	var name string
	err = row.Scan(&title, &name)
	if err != nil {
		return
	}
	return fmt.Sprintf("%s - %s", name, title), err
}

// Get is
func (h Books) Get(id string) (b *Book, err error) {
	row := h.conn.QueryRow(context.Background(), `select b.id, b.title, b.original_title, 
	b.original_language, b.writing_start, b.writing_end, b.publication, b.annotation, b.image, 
	a.id, a.name, a.original_name
	from books b 
	left join authors a on b.author=a.id
	where NOT b.deleted and b.id = $1`, id)
	b = &Book{}
	a := &Author{}
	err = row.Scan(&b.ID, &b.Title, &b.OriginalTitle, &b.OriginalLanguage, &b.WritingStart,
		&b.WritingEnd, &b.Publication, &b.Annotation, &b.Image, &a.ID, &a.Name, &a.OriginalName)
	if err != nil {
		return
	}
	b.Author = a
	rows, err := h.conn.Query(context.Background(), `select g.id, g.name 
		from book_genres bg
		inner join genres g on bg.genre_id=g.id
		where bg.book_id=$1;`, b.ID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		g := &Genre{}
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			return
		}
		b.Genres = append(b.Genres, g)
	}
	b.Author = a
	return
}

// Init is
func (h Books) Init() (err error) {
	log.Println("initializing books.")
	err = h.AddList(bList)
	return
}

// Remove is
func (h Books) Remove(id string) (err error) {
	_, err = h.conn.Exec(context.Background(), `update books set deleted=true where id=$1`, id)
	return
}

// Search is
func (h Books) Search(value string) (bList []*BookPreview, err error) {
	value = "%" + strings.ToLower(value) + "%"
	rows, err := h.conn.Query(context.Background(), `select b.id, b.title, b.original_title, b.publication, 
	a.name
	from books b 
	left join authors a on b.author=a.id
	where NOT b.deleted AND
	(lower(b.title) LIKE $1 OR lower(b.original_title) LIKE $1 OR lower(a.name) LIKE $1)
	order by b.title  
	limit $2`, value, searchNumber)
	if err != nil {
		return
	}
	defer rows.Close()
	var b *BookPreview
	var t time.Time
	for rows.Next() {
		b = &BookPreview{}

		err = rows.Scan(&b.ID, &b.Title, &b.OriginalTitle, &t, &b.Author)
		if err != nil {
			return
		}
		splitted := strings.Split(b.Author, " ")
		lastName := len(splitted) - 1
		b.Author = ""
		for _, v := range splitted[:lastName] {
			r := []rune(v)
			b.Author += fmt.Sprintf("%s. ", string(r[0]))
		}
		b.Author += splitted[lastName]
		b.Publication = t.Year()
		bList = append(bList, b)
	}
	return
}
