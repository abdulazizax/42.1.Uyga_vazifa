package main

import (
	"fmt"
	"log"
	"net/http"
	"user/api"
	"user/storage"
	"user/storage/db"
)

func main() {
	db := db.DBConnect()
	user := storage.NewUser(db)
	r := api.Router(user)
	fmt.Println("server is listening on 9001")
	err := http.ListenAndServe(":9001", r)
	if err != nil {
		log.Fatal(err)
	}
}
