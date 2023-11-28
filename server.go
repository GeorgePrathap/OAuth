package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "golang.org/x/oauth2"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", handleHome)
	r.HandleFunc("/google/login", handleGoogleLogin)
	r.HandleFunc("/facebook/login", handleFacebookLogin)

	http.Handle("/", r)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	file := "login.html"
	http.ServeFile(w, r, file)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside google login")
}

func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside facebook login")
}
