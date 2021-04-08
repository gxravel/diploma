package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
)

// User is the model of the databse table User
type User struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	AdminRights bool   `json:"admin,boolean"`
}

// Users is
type Users struct {
	conn *pgxpool.Pool
}

// UserMethods is
type UserMethods interface {
	Add(*User) (string, error)
	CheckToken(string) (bool, error)
	ClearToken(string) error
	CreateTable() error
	DropTable() error
	GetID(string) (string, error)
	GetLogin(string) (string, error)
	GetPassword(string) (string, bool, error)
	Edit(*User) error
	GetList() ([]*User, error)
	Init() error
	IsAdmin(string) (bool, error)
	UpdateToken(string, string) error
}

// Add inserts into User login, password and admin
func (h Users) Add(user *User) (token string, err error) {
	result, err := h.conn.Exec(context.Background(),
		`INSERT INTO users (login, password, admin) VALUES ($1, $2, $3) returning token;`,
		user.Login, user.Password, user.AdminRights)
	if err != nil {
		return
	}
	token = result.String()
	return
}

// Edit is
func (h Users) Edit(u *User) (err error) {
	_, err = h.conn.Exec(context.Background(), `update books set login=$1, password=$2, 
	admin=$3, token=$4 where id=$5;`, u.Login, u.Password, u.AdminRights, u.Token, u.ID)
	if err != nil {
		return
	}
	return
}

// GetList is
func (h Users) GetList() (list []*User, err error) {
	var nullString sql.NullString
	rows, err := h.conn.Query(context.Background(), `select id, login, password, admin, token 
	from users`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		u := &User{}
		err = rows.Scan(&u.ID, &u.Login, &u.Password, &u.AdminRights, &nullString)
		if err != nil {
			return
		}
		if nullString.Valid {
			u.Token = nullString.String
		} else {
			u.Token = ""
		}
		list = append(list, u)
	}
	return
}

// ClearToken updates user to set token as "" (empty string)
func (h Users) ClearToken(token string) (err error) {
	_, err = h.conn.Exec(context.Background(), `UPDATE users SET token=NULL WHERE token=$1;`, token)

	return
}

// CreateTable is
func (h Users) CreateTable() (err error) {
	_, err = h.conn.Exec(context.Background(), `Create EXTENSION IF NOT EXISTS "uuid-ossp"`)
	if err != nil {
		return
	}
	_, err = h.conn.Exec(context.Background(), `
	create table users (
		id uuid primary key default uuid_generate_v4(),
		login varchar(20) unique,
		password varchar(20) NOT NULL,
		admin bool default 'f',
		token uuid unique default uuid_generate_v4() 
		);`)
	return
}

// DropTable is
func (h Users) DropTable() (err error) {
	_, err = h.conn.Exec(context.Background(),
		`drop table if exists users cascade;`)
	return
}

// GetID finds id by token
func (h Users) GetID(token string) (id string, err error) {
	row := h.conn.QueryRow(context.Background(), `SELECT id FROM users WHERE token=$1;`, token)
	err = row.Scan(&id)
	return
}

// GetLogin finds login by token
func (h Users) GetLogin(token string) (login string, err error) {
	row := h.conn.QueryRow(context.Background(), `SELECT login FROM users WHERE token=$1;`, token)
	err = row.Scan(&login)
	return
}

// GetPassword finds password by login
func (h Users) GetPassword(login string) (password string, admin bool, err error) {
	row := h.conn.QueryRow(context.Background(), `SELECT password, admin FROM users WHERE login=$1;`, login)
	err = row.Scan(&password, &admin)
	return
}

// CheckToken is
func (h Users) CheckToken(token string) (result bool, err error) {
	rows, err := h.conn.Query(context.Background(), `select token from users where token=$1;`, token)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	defer rows.Close()
	return
}

// Init is
func (h Users) Init() (err error) {
	log.Println("initializing users.")
	_, err = h.conn.Exec(context.Background(), `insert into users (login, password, admin) 
	values ($1,$2,$3);`, "ravel123", "rav2647A", true)
	return
}

// IsAdmin checks if User.login has admin rights
func (h Users) IsAdmin(token string) (admin bool, err error) {
	row := h.conn.QueryRow(context.Background(), `SELECT admin FROM users WHERE token=$1;`, token)
	err = row.Scan(&admin)
	return
}

// UpdateToken updates User with provided login to set new token
func (h Users) UpdateToken(login string, token string) (err error) {
	_, err = h.conn.Exec(context.Background(), `UPDATE users SET token=$1 WHERE login=$2;`, token, login)
	return
}
