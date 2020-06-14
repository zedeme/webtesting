package routers

import (
	"encoding/json"
	"net/http"

	"github.com/zedeme/webtesting/db"
	"github.com/zedeme/webtesting/models"
)

//Register is the function to create DB register for the user
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in the dates "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The email is required ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password need 6 characters to continue ", 400)
		return
	}

	_, founded, _ := db.UserAlreadyExist(t.Email)

	if founded {
		http.Error(w, "Email already used with other user", 400)
		return
	}
	_, status, err := db.InsertCredentials(t)
	if err != nil {
		http.Error(w, "Error with user registration: " + err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Error with insert credentials.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
