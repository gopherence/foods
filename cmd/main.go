package main

import (
	"gopherence/web/app/controller"
	"log"
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/foods", controller.FoodHandle)

	// Static file
	fs := http.FileServer(http.Dir("./www"))
	http.Handle("/", fs)

	// Start http server
	addr := ":8080"
	log.Printf("Listening on http://%s", addr)
	http.ListenAndServe(addr, nil)
}
