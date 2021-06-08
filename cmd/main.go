package main

import (
	"github.com/gopherence/foods/app/controller"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/foods", controller.FoodHandle)
	//http.HandleFunc("/instance", controller.InstanceHandle)

	fs := http.FileServer(http.Dir("./var/www"))
	http.Handle("/", fs)

	addr := ":8080"
	log.Printf("Listening on http://%s", addr)
	log.Printf("Architecture: %s | OS: %s", runtime.GOARCH, runtime.GOOS)
	http.ListenAndServe(addr, nil)
}
