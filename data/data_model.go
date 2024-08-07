package data

import (
	"fmt"
	"os"
)

type Data interface {
	GetById(int64) interface{}
	Create()
	Save()
}

var pgxConnectionString string

func init() {
	pgxConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s timezone=UTC connect_timeout=5 password=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_SSL_MODE"),
	)
}
