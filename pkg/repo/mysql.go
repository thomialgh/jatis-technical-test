package repo

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type Option struct {
	Username string
	Password string
	Host     string
	Port     string
}

func Init(opt Option) {
	var err error
	DB, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
	))

	if err != nil {
		log.Fatalf("Failed to connect database with err : %v", err)
	}
}
