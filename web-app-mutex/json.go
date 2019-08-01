package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		var address Address
		json.NewDecoder(r.Body).Decode(&user)
		json.NewDecoder(r.Body).Decode(&address)

		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
		fmt.Fprintf(w, "The Street is %s and city is %s", address.Street, address.City)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "Pavan",
			Lastname:  "Juttada",
			Age:       35,
		}

		NewYork := Address{
			Street: "Manhattan",
			City:   "NYC",
		}

		json.NewEncoder(w).Encode(peter)
		json.NewEncoder(w).Encode(NewYork)
	})

	http.ListenAndServe(":8088", nil)
}
