package controller

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func InstanceHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[instance] %s %s", r.Method, r.URL.Path)

	responseBody := []byte{}
	statusCode := http.StatusMethodNotAllowed

	switch r.Method {
	case "GET":
		responseBody, statusCode = currentInstance(w)
	}

	w.Header().Set("Content-Type", "text/plain; charset=ascii")
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

func currentInstance(w http.ResponseWriter) ([]byte, int) {
	info := fmt.Sprintf("Architecture: %s - OS: %s", runtime.GOARCH, runtime.GOOS)
	return []byte(info), http.StatusOK
}
