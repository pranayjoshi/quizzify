package user

import (
	// "context"
	"encoding/json"
	"net/http"

	// "encoding/json"
	"fmt"
	// "net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store := NewStore()
	store.Create(&user)

	fmt.Printf("Successfully created user: %v\n", user.Name)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store := NewStore()
	u, err := store.GetByName(user.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if u.Password != user.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": u.Name, "name": u.Name})

	fmt.Printf("Successfully logged in user: %v\n", u.Name)
}

func GetResultsbyName(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}

	store := NewStore()

	results, err := store.GetByName(name + "/results")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

	fmt.Printf("Successfully fetched results for user: %v\n", name)
}
