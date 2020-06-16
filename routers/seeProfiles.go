package routers

import (
	"encoding/json"
	"net/http"

	"github.com/zedeme/webtesting/db"
)

//SeeProfile is the function to see user profile
func SeeProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) == 0 {
		http.Error(w, "We need the ID to continue... its important", http.StatusBadRequest)
		return
	}
	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "We have a problem getting profile "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
