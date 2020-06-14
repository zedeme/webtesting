package middlew

import (
	"net/http"

	"github.com/zedeme/webtesting/db"
)

//CheckDB is the middelware for check DB
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ConnectionCheck() == false {
			http.Error(w, "Houston we have a problem, dont have connection with DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
