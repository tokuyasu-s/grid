package db

import (
	"database/sql"
	"fmt"
	config "grid/conf"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

func init() {
	Db, err := ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()
}

func ConnectDb() (db *sql.DB, err error) {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Config.Host, config.Config.User, config.Config.Password, config.Config.Database)
	Db, err := sql.Open(config.Config.SQLDriver, connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return Db, err
}
