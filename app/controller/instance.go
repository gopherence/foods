package controller

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func InstanceHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[instance] %s %s", r.Method, r.URL.Path)

	switch r.Method {
	case "GET":
		currentInstance(w)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func currentInstance(w http.ResponseWriter) {
	info := fmt.Sprintf("Architecture: %s - OS: %s", runtime.GOARCH, runtime.GOOS)

	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.WriteHeader(http.StatusOK)
	writeJSON(w, info)
}
