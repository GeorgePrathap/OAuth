package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	clientID     = ""
	clientSecret = ""
	redirectURL  = "http://localhost:8080/google/callback"
	scopes       = []string{"https://www.googleapis.com/auth/plus.login"}

	googleConfig = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
	}
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", handleHome)
	r.HandleFunc("/google/login", handleGoogleLogin)
	r.HandleFunc("/google/callback", handleGoogleCallBack)
	r.HandleFunc("/facebook/login", handleFacebookLogin)

	http.Handle("/", r)

	log.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	file := "login.html"
	http.ServeFile(w, r, file)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallBack(w http.ResponseWriter, r *http.Request) {
	log.Println("inside callback ...")

	// Retrieve the authorization code from the query parameters
	code := r.URL.Query().Get("code")
	if code == "" {
		log.Println("Authorization code not found")
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for an access token
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %s\n", err)
		http.Error(w, fmt.Sprintf("Failed to exchange token: %s", err), http.StatusInternalServerError)
		return
	}

	// Now you have the access token, and you can use it to make API requests.
	// For example, you can print the user's profile information.
	client := googleConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/plus/v1/people/me")
	if err != nil {
		log.Printf("Failed to get user profile: %s\n", err)
		http.Error(w, fmt.Sprintf("Failed to get user profile: %s", err), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var profile map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		log.Printf("Failed to parse JSON response: %s\n", err)
		http.Error(w, fmt.Sprintf("Failed to parse JSON response: %s", err), http.StatusInternalServerError)
		return
	}

	log.Println("working")

	for key, value := range profile {
		log.Printf("%v : %v\n", key, value)
	}
}

func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside facebook login")
}
