package main

import (
	"log"

	"github.com/zedeme/webtesting/db"
	"github.com/zedeme/webtesting/handlers"
)

func main() {
	if db.ConnectionCheck() == false {
		log.Fatal("Houston, Houston, we have a problem...")
	}
	handlers.Handlers()
}
