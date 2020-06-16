package routers

import (
	"encoding/json"
	"net/http"

	"github.com/zedeme/webtesting/db"
	"github.com/zedeme/webtesting/jwt"
	"github.com/zedeme/webtesting/models"
)

//Login is the function for login users
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Username or Password are invalid "+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "Email is requierd ", 400)
		return
	}
	user, exist := db.TryLogin(u.Email, u.Password)

	if exist == false {
		http.Error(w, "Username or Password are invalid ", 400)
		return
	}

	jwtKey, err := jwt.GeneratorJWT(user)
	if err != nil {
		http.Error(w, "Username or Password are invalid "+err.Error(), 400)
		return
	}

	resp := models.LoginResp{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey
		Expires: expirationTime

	})*/
}
