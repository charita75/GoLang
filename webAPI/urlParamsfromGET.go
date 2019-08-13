package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type tag struct {
	param1, param2, param3 string
}

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	var t tag

	t.param1 = req.FormValue("name")
	t.param2 = req.FormValue("phone")
	t.param3 = req.FormValue("age")
	fmt.Println(t)

}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		queryParamDisplayHandler(res, req)
	})
	println("Enter this in your browser:  http://localhost:8089/?name=charita&phone=123-1234")
	http.ListenAndServe(":8082", nil)

	db, err := sql.Open("mysql", "root:Vijju@123@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO user VALUES ( ?,?,? )")

	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}
