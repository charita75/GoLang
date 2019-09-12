package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

//PageVariables to Store Date and Time
type PageVariables struct {
	Date string
	Time string
}

func main() {

	http.HandleFunc("/", HomePage)
	//http.ListenAndServe(":8081", nil)
	log.Fatal(http.ListenAndServe(":8083", nil))

}

//HomePage Function
func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	HomePageVariables := PageVariables{
		Date: now.Format("01-02-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("homepage.html")
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, HomePageVariables)

}
