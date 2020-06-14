package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/zedeme/webtesting/middlew"
	"github.com/zedeme/webtesting/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handlers are the all handlers have in own website
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
