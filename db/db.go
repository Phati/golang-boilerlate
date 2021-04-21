package db

import (
	"fmt"

	"github.com/Phati/golang-boilerplate/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB *sqlx.DB
)

func InitDB() {
	db, err := sqlx.Connect(config.GetDBConfig().DBDriver, config.GetDBConfig().DBuri)
	if err != nil {
		panic(fmt.Sprintf("db: could not connect to database, %s", err.Error()))
	}
	fmt.Println("Connected to database")
	DB = db
}
