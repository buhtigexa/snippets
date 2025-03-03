package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Config config

type config struct {
	db *sql.DB
}

//func init() {
//	user := os.Getenv("USER")
//	pwd := os.Getenv("PASSWORD")
//	db := os.Getenv("DB")
//	Config.db = openDB(user, pwd, db)
//}

func openDB(user, pwd, dbname string) *sql.DB {
	dsn := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable", user, pwd, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
