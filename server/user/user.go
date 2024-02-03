package user

import (
	"net/http"
)

type User struct {
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

}
func LoginUser(w http.ResponseWriter, r *http.Request) {

}
