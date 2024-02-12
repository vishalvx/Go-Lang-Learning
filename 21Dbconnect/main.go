package main

import (
	"log"
	"net/http"

	"github.com/vishalvx/dbconnect/router"
)

func main() {
	log.Print("--- MONGO DB SERVER ---")

	router := router.Router()

	log.Print("Listening on http://localhost:3000 ...")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal("Fail to Start Server!!")
	}

}
