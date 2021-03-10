package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// User - struct for parsing JSON
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func hendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		user := User{}
		json.Unmarshal([]byte(b), &user)

		fmt.Println("Server accept post request, body -> ", user)
	}
}

func main() {

	http.HandleFunc("/", hendler)

	fmt.Printf("Starting server for testing HTTP POST...\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
