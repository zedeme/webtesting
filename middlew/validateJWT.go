package middlew

import(
	"net/http"
	"github.com/zedeme/webtesting/routers"
)

//ValidateJWT is the middlew validate for JWT
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_,_,_ err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "JWT dont valid" + err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}