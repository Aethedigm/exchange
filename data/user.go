package data

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

type User struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) GetById(id int64) *User {
	conn, err := pgx.Connect(context.Background(), pgxConnectionString)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close(context.Background())

	usr := new(User)
	err = conn.QueryRow(context.Background(), "select * from users where id=$1", id).Scan(&usr)
	if err != nil {
		log.Panic(err)
	}

	return usr
}

func (u *User) Create() {
	panic("not implemented")
}

func (u *User) Save() {
	panic("not implemented")
}
