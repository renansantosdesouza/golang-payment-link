package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/payment-link/config"
)

func DBConnetion() *sql.DB {
	conn := config.GetDBConnection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
