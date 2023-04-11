package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	age       int    `json:"age"`
	die       bool
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			return
		}

		_, err = fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.age)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			age:       25,
			die:       true,
		}

		err := json.NewEncoder(w).Encode(peter)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
