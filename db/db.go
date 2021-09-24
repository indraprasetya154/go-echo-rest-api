package db

import (
	"database/sql"
	"echo-rest/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	db, err = sql.Open("mysql", conf.DB_USERNAME+":"+conf.DB_PASSWORD+"@/"+conf.DB_NAME)
	if err != nil {
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
