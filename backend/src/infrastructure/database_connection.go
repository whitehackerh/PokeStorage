package DB

import "database/sql"

var Db *sql.DB

func ConnectDb() *sql.DB {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=pokestorage password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	return Db
}