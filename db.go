package db

import "database/sql"

//Conn connects to Database
func Conn() (db *sql.DB) {
	
	connString := "root:fatima1234@tcp(localhost:3306)/trial"
	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}
	return db
}
