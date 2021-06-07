package main

import (
	"github.com/gopherence/foods/app/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foods", controller.FoodHandle)

	fs := http.FileServer(http.Dir("./var/www"))
	http.Handle("/", fs)

	addr := ":8080"
	log.Printf("Listening on http://%s", addr)
	http.ListenAndServe(addr, nil)
}
