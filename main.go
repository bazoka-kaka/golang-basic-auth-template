package main

import (
	"fmt"
	"net/http"
)

var (
	USERNAME = "benzion"
	PASSWORD = "yehezkel"
)

func main() {
	http.HandleFunc("/", handleLogin)

	fmt.Println("server running on port 3000")
	http.ListenAndServe(":3000", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("something went wrong"))
		return
	}

	isValid := (username == USERNAME) && (password == PASSWORD)

	if !isValid {
		w.Write([]byte("wrong username/password"))
		return
	}

	w.Write([]byte("User logged in"))
}
