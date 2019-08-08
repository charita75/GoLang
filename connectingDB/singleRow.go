//Querying a single row from the table

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
 * Tag... - a very simple struct
 */
type Tag1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:Vijju@123@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	var tag Tag1
	err = db.QueryRow("SELECT idtest_t, test_name FROM test_t where idtest_t= ?", 1).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// and then print out the tag's Name attribute
	log.Println(tag.Name)
	log.Println(tag.ID)
}
