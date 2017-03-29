//Package store is an adaptor to the sqlite db that contains the data we need.
//It updates the db based on the data received from poller and provides helper methods for main to query data.
package store

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB


//Connect to db
func Connect(dbPath string) *sql.DB{
	db, err := sql.Open("sqlite3", dbPath)
  if err != nil {
		panic(err)
	}
	return db
}

// func PollWF2(data []byte) {
// 	update := parseWf2(data)

// }
